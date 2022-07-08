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
			WithOnPacket(onPacket(sr)).
			WithOnDisconnected(onDisconnected(sr)).
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
		s := sr.Create(conn, encrypter, decrypter)
		err := sr.Add(s)

		if err != nil {
			log.Printf("Could not add session to registry. Rejecting connection (%s): %s", conn.RemoteAddr(), err)
			return
		}

		// Send client handshake
		writer.WriteHandshake(s)(MAJOR_VERSION, MINOR_VERSION, ivRecv[:], ivSend[:], LOCALE)
	}
}

func onPacket(sr session.SessionRegistry) func(conn net.Conn, data []byte) {
	return func(conn net.Conn, data []byte) {
		s, err := sr.Get(conn.RemoteAddr().String())
		if err != nil {
			log.Printf("Could not find session in session registry (%s)", conn.RemoteAddr())
			return
		}

		p := packet.Packet(s.Decrypt(data))
		log.Printf("Received packet (%s): %s", s.ID(), p)
	}
}

func onDisconnected(sr session.SessionRegistry) func(conn net.Conn, reason error) {
	return func(conn net.Conn, reason error) {
		log.Printf("Client closed connection (%s). Reason: %s", conn.RemoteAddr(), reason)
		sr.Destroy(conn.RemoteAddr().String())
	}
}
