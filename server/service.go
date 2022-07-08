package server

import (
	"context"
	"log"
	"net"
	"sync"

	"github.com/matthieutran/leafre-login/pkg/packet"
	"github.com/matthieutran/leafre-login/pkg/tcp"
	"github.com/matthieutran/leafre-login/server/writer"
)

const (
	MAJOR_VERSION = 95
	MINOR_VERSION = "1"
	LOCALE        = 8
)

func Start(wg *sync.WaitGroup, ctx context.Context) func(host string, port int) {
	return func(host string, port int) {
		err := tcp.NewServer().
			WithOnConnected(onConnected).
			WithOnPacket(onPacket).
			WithOnDisconnected(onDisconnected).
			Start(wg, ctx)(host, port)

		if err != nil {
			log.Fatal("Could not start server: ", err)
		}
	}
}

func onConnected(conn net.Conn) {
	log.Printf("New client connection (%s)", conn.RemoteAddr())

	ivRecv := []byte{0, 0, 0, 0}
	ivSend := []byte{0, 0, 0, 0}

	writer.WriteHandshake(conn)(MAJOR_VERSION, MINOR_VERSION, ivRecv, ivSend, LOCALE)
}

func onPacket(conn net.Conn, data []byte) {
	p := packet.Packet(data)
	log.Println(p)
}

func onDisconnected(conn net.Conn, reason error) {
	log.Printf("Client closed connection (%s). Reason: %s", conn.RemoteAddr(), reason)
}
