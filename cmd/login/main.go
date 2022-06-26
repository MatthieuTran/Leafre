package main

import (
	"crypto/rand"
	"log"
	"sync"

	"github.com/matthieutran/leafre-login/internal/login/net/codec"
	"github.com/matthieutran/leafre-login/internal/login/net/writer"
	"github.com/matthieutran/tcpserve"
)

const (
	VERSION       = 83
	MINOR_VERSION = "1"
	LOCALE        = 8
	PORT          = 8484
)

func main() {
	var wg sync.WaitGroup

	log.Println("Leafre - Login Server")

	logger := func(msg string) {
		log.Println(msg)
	}

	onConnected := func(s *tcpserve.Session) {
		var ivRecv, ivSend [4]byte // IV Keys for the codec
		rand.Read(ivRecv[:])       // Randomize recv key
		rand.Read(ivSend[:])       // Randomize send key

		encrypter, decrypter := codec.GenerateCodecs(ivRecv, ivSend) // Create codec
		s.SetEncrypter(encrypter)
		s.SetDecrypter(decrypter)

		// Send handshake
		data := writer.WriteHandshake(VERSION, MINOR_VERSION, ivRecv, ivSend, LOCALE)
		log.Println(data)
		s.Write(data)
	}

	onPacket := func(s *tcpserve.Session, data []byte) {
		log.Printf("Packet: % X", data)
	}

	wg.Add(1)
	server := tcpserve.NewServer(tcpserve.WithPort(PORT), tcpserve.WithLoggers(logger, nil), tcpserve.WithOnConnected(onConnected), tcpserve.WithOnPacket(onPacket))
	server.Start(wg)

	// Block until all goroutines are done
	wg.Wait()
}
