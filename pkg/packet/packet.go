package packet

import (
	"fmt"
)

// A Packet is a structured byte slice.
type Packet []byte

func NewPacket(buf []byte) Packet {
	return buf
}

// Header returns the header of the packet
func (p Packet) Header() []byte {
	return p[:2]
}

// Bytes returns the packet without the header
func (p Packet) Bytes() []byte {
	return []byte(p[2:])
}

// String implements the fmt.Stringer interface by returning itself as a readable string
func (p Packet) String() string {
	return fmt.Sprintf("[% X] % X", p.Header(), p.Bytes())
}
