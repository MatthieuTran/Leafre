package net

import (
	"crypto/rand"
	"log"
	"sync"

	"github.com/matthieutran/duey"
	"github.com/matthieutran/leafre-login/internal/login/net/codec"
	"github.com/matthieutran/leafre-login/internal/login/net/handler"
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

func onPacket(es *duey.EventStreamer, handlers map[uint16]handler.PacketHandler) func(*tcpserve.Session, []byte) {
	return func(s *tcpserve.Session, data []byte) {
		var p packet.Packet
		p.WriteBytes(data)

		header := p.ReadShort()

		// Check if header has a handler
		if h, ok := handlers[header]; ok {
			log.Printf("Handling %s: [%X] % X\n", h.Name(), header, p)
			h.Handle(s, es, p)
		} else {
			log.Printf("Unhandled Packet: [%X] %s\n", header, p)
		}
	}
}

func InitHandlers() map[uint16]handler.PacketHandler {
	// Create handler collection
	handlers := make(map[uint16]handler.PacketHandler)
	addHandler := func(opcode uint16, h handler.PacketHandler) {
		handlers[opcode] = h
	}

	addHandler(handler.OpCodeCheckPassword, &handler.HandlerCheckPassword{}) // 0x00
	addHandler(handler.OpCodeWorldRequest, &handler.HandlerWorldRequest{})   // 0xB

	return handlers
}

func BuildServer(wg sync.WaitGroup, s *duey.EventStreamer) *tcpserve.Server {
	logger := func(msg string) {
		log.Println(msg)
	}

	handlers := InitHandlers()

	server := tcpserve.NewServer(
		tcpserve.WithPort(PORT),
		tcpserve.WithLoggers(logger, nil),
		tcpserve.WithOnConnected(onConnected),
		tcpserve.WithOnPacket(onPacket(s, handlers)),
	)
	server.Start(wg)

	return server
}
