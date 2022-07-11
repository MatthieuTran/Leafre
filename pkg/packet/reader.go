package packet

import (
	"bytes"
	"encoding/binary"
)

// A PacketReader is used to read structured data from a byte buffer.
//
// A PacketReader implements the io.Reader interfaces by reading from a byte slice.
type PacketReader interface {
	ReadUInt16() uint16         // WriteUInt16 reads data of type uint16 and returns it
	ReadUInt32() uint32         // WriteUInt32 reads data of type uint32 and returns it
	ReadUInt64() uint64         // WriteUInt64 reads data of type uint64 and returns it
	ReadString() string         // WriteString reads data of type string and returns it
	ReadAvailableBytes() []byte // ReadAvailableBytes reads the rest of the byte buffer
	ReadBytes(n int) []byte     // ReadBytes reads `n` bytes into a buffer and returns it
	ReadOne() byte              // ReadOne reads 1 byte and returns it
}

type maplePacketReader struct {
	reader *bytes.Reader
}

func NewPacketReader(buf []byte) PacketReader {
	r := maplePacketReader{}
	r.reader = bytes.NewReader(buf)

	return &r
}

func (p maplePacketReader) ReadBytes(n int) (buf []byte) {
	buf = make([]byte, n)
	binary.Read(p.reader, binary.LittleEndian, &buf)
	return
}

func (p maplePacketReader) ReadAvailableBytes() (buf []byte) {
	return p.ReadBytes(p.reader.Len())
}

func (p maplePacketReader) ReadOne() (b byte) {
	binary.Read(p.reader, binary.LittleEndian, &b)
	return
}

func (p maplePacketReader) ReadUInt16() (n uint16) {
	binary.Read(p.reader, binary.LittleEndian, &n)
	return
}

func (p maplePacketReader) ReadUInt32() (n uint32) {
	binary.Read(p.reader, binary.LittleEndian, &n)
	return
}

func (p maplePacketReader) ReadUInt64() (n uint64) {
	binary.Read(p.reader, binary.LittleEndian, &n)
	return
}

func (p maplePacketReader) ReadString() (s string) {
	size := p.ReadUInt16()
	buf := p.ReadBytes(int(size))
	s = string(buf)

	return
}
