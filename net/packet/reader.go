package packet

import (
	"bytes"
	"encoding/binary"
)

// A PacketReader is used to read structured data from a byte buffer.
//
// A PacketReader implements the io.Reader interfaces by reading from a byte slice.
type PacketReader interface {
	ReadUInt16(*uint16) (err error)    // WriteUInt16 reads data of type uint16 into the given parameter
	ReadUInt32(*uint32) (err error)    // WriteUInt32 reads data of type uint32 into the given parameter
	ReadUInt64(*uint64) (err error)    // WriteUInt64 reads data of type uint64 into the given parameter
	ReadString() (s string, err error) // WriteString reads data of type string into the given parameter
	Read(p []byte) (n int, err error)  // PacketReader implements `io.Reader`, appending `len(p)` bytes from `p` to the packet
	ReadByte() (byte, error)           // PacketReader implements `io.ByteReader`, reading and returning the next byte from the packet
}

type maplePacketReader struct {
	reader *bytes.Reader
}

func NewPacketReader(buf []byte) PacketReader {
	r := maplePacketReader{}
	r.reader = bytes.NewReader(buf)

	return &r
}

func (p maplePacketReader) Read(buf []byte) (n int, err error) {
	return n, binary.Read(p.reader, binary.LittleEndian, buf)
}

func (p maplePacketReader) ReadByte() (b byte, err error) {
	return b, binary.Read(p.reader, binary.LittleEndian, &b)
}

func (p maplePacketReader) ReadUInt16(n *uint16) (err error) {
	return binary.Read(p.reader, binary.LittleEndian, n)
}

func (p maplePacketReader) ReadUInt32(n *uint32) (err error) {
	return binary.Read(p.reader, binary.LittleEndian, n)
}

func (p maplePacketReader) ReadUInt64(n *uint64) (err error) {
	return binary.Read(p.reader, binary.LittleEndian, n)
}

func (p maplePacketReader) ReadString() (s string, err error) {
	size, err := p.ReadByte()
	if err != nil {
		return
	}

	buf := make([]byte, size)
	_, err = p.Read(buf)
	s = string(buf)

	return
}
