package net

import (
	"io"
	"log"
	"net"

	"github.com/matthieutran/leafre-login/net/packet"
	"github.com/matthieutran/leafre-login/net/writer"
)

const (
	VERSION               = 95  // majorVersion
	MINOR_VERSION         = "1" // minorVersion
	LOCALE                = 8   // locale GMS = 8
	ENCRYPTED_HEADER_SIZE = 4   // packet's header size
)

// ReadPacketSize decodes the packet length (of the rest of the packet) from the encrypted header.
func ReadPacketSize(r io.Reader) (int, error) {
	buf := make([]byte, ENCRYPTED_HEADER_SIZE)
	_, err := r.Read(buf)
	if err != nil {
		return 0, err
	}

	x := uint16(buf[0]) + uint16(buf[1])*0x100
	y := uint16(buf[2]) + uint16(buf[3])*0x100

	return int(x ^ y), nil
}

// HandleConn blocks until a new packet comes in and handles it accordingly
func HandleConn(conn net.Conn) {
	log.Printf("New connection (%s)", conn.RemoteAddr())

	// Ensure connection is closed at end
	defer conn.Close()

	// Send Handshake to client
	writer.WriteHandshake(conn, VERSION, MINOR_VERSION, []byte{0, 0, 0, 0}, []byte{0, 0, 0, 0}, LOCALE)

	// Handle incoming packets
	for {
		n, err := ReadPacketSize(conn)
		if n == 0 || err != nil {
			// EOF
			log.Printf("Client closed connection (%s). Error: %s", conn.RemoteAddr(), err)
			break
		}
		log.Println("Incoming packet of size:", n)

		// Read Packet
		buf := make([]byte, n)
		n, err = conn.Read(buf)
		if n == 0 || err != nil {
			// EOF
			log.Printf("Client closed connection (%s). Error: %s", conn.RemoteAddr(), err)
			break
		}
		log.Println("Packet:", packet.Packet(buf))
	}
}
