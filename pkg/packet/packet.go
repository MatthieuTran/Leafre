package packet

import (
	"encoding/binary"
	"fmt"
)

// A Packet is a structured byte slice.
type Packet []byte

func NewPacket(buf []byte) Packet {
	return buf
}

// Header returns the header of the packet
func (p Packet) Header() uint16 {
	return binary.LittleEndian.Uint16(p[:2])
}

// Bytes returns the packet without the header
func (p Packet) Bytes() []byte {
	return []byte(p[2:])
}

func (p Packet) bytes() []byte {
	return []byte(p)
}

// String implements the fmt.Stringer interface by returning itself as a readable string
func (p Packet) String() string {
	return fmt.Sprintf("% X", p.bytes())
}
