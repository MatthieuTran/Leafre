package tcp

import (
	"io"
	"net"
)

const (
	ENCRYPTED_HEADER_SIZE = 4 // packet's header size
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

// handleConn blocks until a new packet comes in and handles it accordingly
func (s *Server) handleConn(conn net.Conn) {
	var reason error
	s.onConnected(conn)

	// Ensure connection is closed at end
	defer func() {
		s.onDisconnected(conn, reason)
		conn.Close()
	}()

	// Handle incoming packets
	for {
		n, err := ReadPacketSize(conn)
		if n == 0 || err != nil {
			// EOF
			reason = err
			break
		}
		// Read Packet
		buf := make([]byte, n)
		n, err = conn.Read(buf)
		if n == 0 || err != nil {
			// EOF
			reason = err
			break
		}
		s.onPacket(conn, buf)
	}
}
