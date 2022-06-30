package net

import (
	"crypto/rand"
	"log"
	"sync"

	"github.com/matthieutran/duey"
	"github.com/matthieutran/leafre-login/internal/login/net/codec"
	"github.com/matthieutran/leafre-login/internal/login/net/handler/auth"
	"github.com/matthieutran/leafre-login/internal/login/net/writer"
	"github.com/matthieutran/packet"
	"github.com/matthieutran/tcpserve"
)

const (
	VERSION       = 95
	MINOR_VERSION = "1"
	LOCALE        = 8
	PORT          = 8484
)

func onConnected(s *tcpserve.Session) {
	var ivRecv, ivSend [4]byte // IV Keys for the codec
	rand.Read(ivRecv[:])       // Randomize recv key
	rand.Read(ivSend[:])       // Randomize send key

	encrypter, decrypter := codec.GenerateCodecs(VERSION, ivRecv, ivSend) // Create codec
	s.SetEncrypter(encrypter)
	s.SetDecrypter(decrypter)

	// Send handshake
	handshakePacket := writer.WriteHandshake(VERSION, MINOR_VERSION, ivRecv, ivSend, LOCALE)
	s.WriteRaw(handshakePacket)
}

func onPacket(es *duey.EventStreamer) func(*tcpserve.Session, []byte) {
	return func(s *tcpserve.Session, data []byte) {
		var p packet.Packet
		p.WriteBytes(data)

		header := p.ReadShort()
		switch header {
		case 0x01: // LOGIN_PASSWORD
			auth.HandleLogin(s, es, p)
		case 0x1A: // EXCEPTION_LOG
			_, msg := p.ReadString()
			log.Println("Received exception log from client:", msg)
		default:
			log.Printf("Unhandled Packet (Header: % X): %s", header, p)
		}

	}
}

func BuildServer(wg sync.WaitGroup, s *duey.EventStreamer) *tcpserve.Server {
	logger := func(msg string) {
		log.Println(msg)
	}

	server := tcpserve.NewServer(
		tcpserve.WithPort(PORT),
		tcpserve.WithLoggers(logger, nil),
		tcpserve.WithOnConnected(onConnected),
		tcpserve.WithOnPacket(onPacket(s)),
	)
	server.Start(wg)

	return server
}
