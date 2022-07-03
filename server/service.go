package server

import (
	"crypto/rand"
	"log"
	"sync"

	"github.com/matthieutran/duey"
	"github.com/matthieutran/leafre-login/codec"
	"github.com/matthieutran/leafre-login/repository"
	"github.com/matthieutran/leafre-login/server/handler"
	"github.com/matthieutran/packet"
	"github.com/matthieutran/tcpserve"
)

const (
	VERSION       = 95
	MINOR_VERSION = "1"
	LOCALE        = 8
	PORT          = 8484
)

func formHandshake(majorVersion uint16, minorVersion string, ivRecv, ivSend [4]byte, locale byte) []byte {
	p := packet.Packet{}
	p.WriteShort(14)            // Length of packet
	p.WriteShort(majorVersion)  // Maple Version (83)
	p.WriteString(minorVersion) // Subversion (1)
	p.WriteBytes(ivRecv[:])     // Recv IV
	p.WriteBytes(ivSend[:])     // Send IV
	p.WriteByte(locale)         // Maple Locale (8)

	return p.Bytes()
}

func onConnected(s *tcpserve.Session) {
	var ivRecv, ivSend [4]byte // IV Keys for the codec
	rand.Read(ivRecv[:])       // Randomize recv key
	rand.Read(ivSend[:])       // Randomize send key

	encrypter, decrypter := codec.GenerateCodecs(VERSION, ivRecv, ivSend) // Create codec
	s.SetEncrypter(encrypter)
	s.SetDecrypter(decrypter)

	// Send handshake
	handshakePacket := formHandshake(VERSION, MINOR_VERSION, ivRecv, ivSend, LOCALE)
	s.WriteRaw(handshakePacket)
}

func onPacket(es *duey.EventStreamer, handlers map[uint16]handler.PacketHandler) func(*tcpserve.Session, []byte) {
	return func(s *tcpserve.Session, data []byte) {
		var p packet.Packet
		p.WriteBytes(data)

		header := p.ReadShort()

		// Check if header has a handler
		if h, ok := handlers[header]; ok {
			log.Printf("Handling %s: [%X] %s\n", h, header, p)
			h.Handle(s, es, p)
		} else {
			log.Printf("Unhandled Packet: [%X] %s\n", header, p)
		}
	}
}

func InitHandlers(es *duey.EventStreamer) map[uint16]handler.PacketHandler {
	// Create handler collection
	handlers := make(map[uint16]handler.PacketHandler)
	addHandler := func(opcode uint16, h handler.PacketHandler) {
		handlers[opcode] = h
	}

	// Inject event streamer into the repository
	userRepository := repository.NewUserRepository(es)

	// Inject repository dependencies into handlers
	handlerCheckPassword := handler.NewHandlerCheckPassword(userRepository)

	addHandler(handler.OpCodeCheckPassword, &handlerCheckPassword)         // 0x00
	addHandler(handler.OpCodeWorldRequest, &handler.HandlerWorldRequest{}) // 0xB

	return handlers
}

func BuildServer(wg *sync.WaitGroup, es *duey.EventStreamer) *tcpserve.Server {
	logger := func(msg string) {
		log.Println(msg)
	}

	// Initialize handlers
	handlers := InitHandlers(es) // the event streamer is injected here

	server := tcpserve.NewServer(
		tcpserve.WithPort(PORT),
		tcpserve.WithLoggers(logger, nil),
		tcpserve.WithOnConnected(onConnected),
		tcpserve.WithOnPacket(onPacket(es, handlers)),
	)
	server.Start(wg)

	return server
}
