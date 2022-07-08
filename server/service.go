package server

import (
	"context"
	"crypto/rand"
	"log"
	"net"
	"sync"

	"github.com/matthieutran/leafre-login/pkg/packet"
	"github.com/matthieutran/leafre-login/pkg/tcp"
	"github.com/matthieutran/leafre-login/server/session"
	"github.com/matthieutran/leafre-login/server/writer"
)

const (
	MAJOR_VERSION = 95
	MINOR_VERSION = "1"
	LOCALE        = 8
)

func Start(wg *sync.WaitGroup, ctx context.Context, sr session.SessionRegistry) func(host string, port int) {
	return func(host string, port int) {
		err := tcp.NewServer().
			WithOnConnected(onConnected(sr)).
			WithOnPacket(onPacket).
			WithOnDisconnected(onDisconnected).
			Start(wg, ctx)(host, port)

		if err != nil {
			log.Fatal("Could not start server: ", err)
		}
	}
}

func onConnected(sr session.SessionRegistry) func(conn net.Conn) {
	return func(conn net.Conn) {
		log.Printf("New client connection (%s)", conn.RemoteAddr())

		// Create IVs
		var ivRecv, ivSend [4]byte
		rand.Read(ivRecv[:])
		rand.Read(ivSend[:])

		// Generate Codecs
		encrypter, decrypter := generateCodecs(MAJOR_VERSION, ivRecv, ivSend)

		// Create Session
		s, err := sr.Create(conn, encrypter, decrypter)
		if err != nil {
			log.Printf("Could not create session. Rejecting connection (%s): %s", conn.RemoteAddr(), err)
			return
		}

		// Send client handshake
		writer.WriteHandshake(s)(MAJOR_VERSION, MINOR_VERSION, ivRecv[:], ivSend[:], LOCALE)
	}
}

func onPacket(conn net.Conn, data []byte) {
	p := packet.Packet(data)
	log.Println(p)
}

func onDisconnected(conn net.Conn, reason error) {
	log.Printf("Client closed connection (%s). Reason: %s", conn.RemoteAddr(), reason)
}
