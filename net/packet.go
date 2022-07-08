package net

import "github.com/matthieutran/leafre-login/pkg/packet"

type PacketWriter interface {
	WriteUInt16(n uint16) (err error)        // WriteUInt16 appends a number of type uint16 to the packet
	WriteUInt32(n uint32) (err error)        // WriteUInt32 appends a number of type uint32 to the packet
	WriteUInt64(n uint64) (err error)        // WriteUInt64 appends a number of type uint64 to the packet
	WriteString(s string) (n int, err error) // WriteString appends a string to the packet
	Packet() packet.Packet                   // Packet returns the packet associated with the writer
	Write(p []byte) (n int, err error)       // PacketWriter implements `io.Writer`, appending `len(p)` bytes from `p` to the packet
	WriteByte(c byte) error                  // PacketWriter implements `io.ByteWriter`, appending byte `c` to the packet
}
