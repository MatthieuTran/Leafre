package ports

import (
	"context"
	"crypto/rand"
	"log"
	"net"
	"sync"

	"github.com/matthieutran/leafre-login/internal/app/handling/writer"
	"github.com/matthieutran/leafre-login/internal/domain/session"
	"github.com/matthieutran/leafre-login/pkg/crypto"
	"github.com/matthieutran/leafre-login/pkg/packet"
	"github.com/matthieutran/leafre-login/pkg/tcp"
)

const (
	MAJOR_VERSION = 95
	MINOR_VERSION = "1"
	LOCALE        = 8
)

func StartTCPServer(wg *sync.WaitGroup, ctx context.Context, ss session.SessionService) func(host string, port int) {
	return func(host string, port int) {
		err := tcp.NewServer().
			WithOnConnected(onConnected(ss)).
			WithOnPacket(onPacket(ss)).
			WithOnDisconnected(onDisconnected(ss)).
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

func onPacket(ss session.SessionService) func(conn net.Conn, data []byte) {
	return func(conn net.Conn, data []byte) {
		id := conn.RemoteAddr().String()
		decrypted, err := ss.DecryptPacket(context.Background(), id, data)
		if err != nil {
			log.Println("Error decrypting packet:", err)
			return
		}

		p := packet.Packet(decrypted)
		log.Printf("Received packet (%s): %s", id, p)
	}
}

func onDisconnected(ss session.SessionService) func(conn net.Conn, reason error) {
	return func(conn net.Conn, reason error) {
		log.Printf("Client closed connection (%s). Reason: %s", conn.RemoteAddr(), reason)
		ss.RemoveSession(context.Background(), conn.RemoteAddr().String())
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
