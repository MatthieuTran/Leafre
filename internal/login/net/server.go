package net

import (
	"crypto/rand"
	"log"
	"sync"

	"github.com/matthieutran/leafre-login/internal/login/net/codec"
	"github.com/matthieutran/leafre-login/internal/login/net/writer"
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
	data := writer.WriteHandshake(VERSION, MINOR_VERSION, ivRecv, ivSend, LOCALE)
	log.Println(data)
	s.Write(data)
}

func onPacket(s *tcpserve.Session, data []byte) {
	log.Printf("Packet: % X", data)
}

func BuildServer(wg sync.WaitGroup) *tcpserve.Server {
	logger := func(msg string) {
		log.Println(msg)
	}

	server := tcpserve.NewServer(tcpserve.WithPort(PORT), tcpserve.WithLoggers(logger, nil), tcpserve.WithOnConnected(onConnected), tcpserve.WithOnPacket(onPacket))
	server.Start(wg)

	return server
}
