package packet

import (
	"bytes"
	"encoding/binary"
	"io"
)

// A PacketWriter is used to write structured data into a byte buffer.
//
// A PacketReader implements the io.Reader interfaces by writing into a byte slice.
type PacketWriter interface {
	WriteUInt16(n uint16) (err error)        // WriteUInt16 appends a number of type uint16 to the packet
	WriteUInt32(n uint32) (err error)        // WriteUInt32 appends a number of type uint32 to the packet
	WriteUInt64(n uint64) (err error)        // WriteUInt64 appends a number of type uint64 to the packet
	WriteString(s string) (n int, err error) // WriteString appends a string to the packet
	Packet() Packet                          // Packet returns the packet associated with the writer
	Write(p []byte) (n int, err error)       // PacketWriter implements `io.Writer`, appending `len(p)` bytes from `p` to the packet
	WriteByte(c byte) error                  // PacketWriter implements `io.ByteWriter`, appending byte `c` to the packet
}

type maplePacketWriter struct {
	buf bytes.Buffer
}

func NewPacketWriter() PacketWriter {
	return &maplePacketWriter{}
}

func (p *maplePacketWriter) Packet() Packet {
	return p.buf.Bytes()
}

func (p *maplePacketWriter) Write(buf []byte) (n int, err error) {
	return n, binary.Write(io.Writer(&p.buf), binary.LittleEndian, buf)
}

func (p *maplePacketWriter) WriteByte(b byte) (err error) {
	return p.buf.WriteByte(b)
}

func (p *maplePacketWriter) WriteUInt16(n uint16) (err error) {
	return binary.Write(io.Writer(&p.buf), binary.LittleEndian, n)
}

func (p *maplePacketWriter) WriteUInt32(n uint32) (err error) {
	return binary.Write(io.Writer(&p.buf), binary.LittleEndian, n)
}

func (p *maplePacketWriter) WriteUInt64(n uint64) (err error) {
	return binary.Write(io.Writer(&p.buf), binary.LittleEndian, n)
}

// WriteString appends a byte containing the size of the string and then the string as a byte slice
func (p *maplePacketWriter) WriteString(s string) (n int, err error) {
	p.WriteByte(byte(len(s)))
	return p.Write([]byte(s))
}
