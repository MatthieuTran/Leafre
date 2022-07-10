package ports

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/binary"
	"log"
	"net"
	"sync"

	"github.com/matthieutran/leafre-login/internal/app/handler"
	"github.com/matthieutran/leafre-login/internal/app/handler/writer"
	"github.com/matthieutran/leafre-login/internal/domain/session"
	"github.com/matthieutran/leafre-login/internal/service"
	"github.com/matthieutran/leafre-login/pkg/crypto"
	"github.com/matthieutran/leafre-login/pkg/packet"
	"github.com/matthieutran/leafre-login/pkg/tcp"
)

const (
	MAJOR_VERSION = 95
	MINOR_VERSION = "1"
	LOCALE        = 8
)

func StartTCPServer(wg *sync.WaitGroup, ctx context.Context, app *service.Application) func(host string, port int) {
	return func(host string, port int) {
		err := tcp.NewServer().
			WithOnConnected(onConnected(app.SessionService)).
			WithOnPacket(onPacket(app.SessionCommunicationService, app.Handlers)).
			WithOnDisconnected(onDisconnected(app.SessionService)).
			Start(wg, ctx)(host, port)

		if err != nil {
			log.Fatal("Could not start server: ", err)
		}
	}
}

func onConnected(ss session.SessionService) func(conn net.Conn) {
	return func(conn net.Conn) {
		log.Printf("New client connection (%s)", conn.RemoteAddr())

		// Create IVs
		var ivRecv, ivSend [4]byte
		rand.Read(ivRecv[:])
		rand.Read(ivSend[:])

		// Generate Codecs
		encrypter, decrypter := generateCodecs(MAJOR_VERSION, ivRecv, ivSend)

		// Create and store this session
		s := session.NewSession(conn, encrypter, decrypter)
		err := ss.CreateSession(context.Background(), s)
		if err != nil {
			log.Printf("Could not add session to registry. Rejecting connection (%s): %s", conn.RemoteAddr(), err)
			return
		}

		// Send client handshake
		writer.WriteHandshake(conn)(MAJOR_VERSION, MINOR_VERSION, ivRecv[:], ivSend[:], LOCALE)
	}
}

type handlersMap map[uint16]handler.PacketHandler

func onPacket(scs session.SessionCommunicationService, handlers handlersMap) func(conn net.Conn, data []byte) {
	return func(conn net.Conn, data []byte) {
		id := conn.RemoteAddr().String()

		// Decrypt incoming packet
		decrypted, err := scs.DecryptPacket(id, data)
		p := packet.Packet(decrypted)
		if err != nil {
			log.Println("Error decrypting packet:", err)
			return
		}

		var header uint16
		r := bytes.NewReader(p.Header())
		binary.Read(r, binary.LittleEndian, &header)

		var buf bytes.Buffer
		if h, ok := handlers[header]; ok {
			// Write packet
			log.Printf("RECV %s: %s\n", h, p)
			h.Handle(&buf, p.Bytes())
		} else {
			log.Printf("RECV Unhandled Packet: %s\n", p)
			return
		}

		// Send packet
		res := buf.Bytes()
		scs.WriteToID(id, res)
		log.Printf("SEND (%s): %s", id, packet.Packet(res))
	}
}

func onDisconnected(ss session.SessionService) func(conn net.Conn, reason error) {
	return func(conn net.Conn, reason error) {
		id := conn.RemoteAddr().String()
		log.Printf("Client closed connection (%s). Reason: %s", conn.RemoteAddr(), reason)
		ss.RemoveSession(context.Background(), id)
	}
}

func generateCodecs(version int, ivRecv, ivSend [4]byte) (encrypter, decrypter func(d []byte) []byte) {
	// Create codecs
	c := crypto.NewCodec(ivRecv, ivSend, version)

	// Create encrypter
	encrypter = func(d []byte) (res []byte) {
		res, _ = c.Encrypt(d, true, true)
		return
	}

	// Create decrypter
	decrypter = func(d []byte) (res []byte) {
		res, _ = c.Decrypt(d, true, true)
		return
	}

	return
}
