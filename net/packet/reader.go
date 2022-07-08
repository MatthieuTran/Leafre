package packet

import "io"

// A PacketReader is used to read structured data from a byte buffer.
//
// A PacketReader implements the io.Reader interfaces by reading from a byte slice.
type PacketReader interface {
	ReadUInt16(n uint16) (err error)
	ReadUInt32(n uint32) (err error)
	ReadUInt64(n uint64) (err error)
	ReadString(s string) (n int, err error)
	Packet() Packet
	io.Reader
	io.ByteReader
}
