package packet

import (
	"bytes"
	"encoding/binary"
)

// A PacketWriter is used to write structured data into a byte buffer.
//
// A PacketWriter implements the io.Writer interfaces by writing into a byte slice.
type PacketWriter interface {
	WriteUInt16(n uint16) (err error)                          // WriteUInt16 appends a number of type uint16 `n` to the packet
	WriteUInt32(n uint32) (err error)                          // WriteUInt32 appends a number of type uint32 `n` to the packet
	WriteUInt64(n uint64) (err error)                          // WriteUInt64 appends a number of type uint64 `n` to the packet
	WriteString(s string) (n int, err error)                   // WriteString appends a string `s` to the packet. The length of the string `s` is appended as a byte before the string.
	WritePaddedString(s string, length int) (n int, err error) // WriteString appends a string `s` of fixed length `length` to the packet. If the string `s` is shorter than `length`, padding is appended to string.
	WriteBytes(p []byte) (n int, err error)                    // WriteBytes appends `len(p)` bytes from `p` to the packet
	WriteOne(c byte) error                                     // WriteOne appends byte `c` to the packet
	Packet() Packet                                            // Packet returns the packet associated with the writer
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

func (p *maplePacketWriter) WriteBytes(buf []byte) (n int, err error) {
	return n, binary.Write(&p.buf, binary.LittleEndian, buf)
}

func (p *maplePacketWriter) WriteOne(b byte) (err error) {
	return p.buf.WriteByte(b)
}

func (p *maplePacketWriter) WriteUInt16(n uint16) (err error) {
	return binary.Write(&p.buf, binary.LittleEndian, n)
}

func (p *maplePacketWriter) WriteUInt32(n uint32) (err error) {
	return binary.Write(&p.buf, binary.LittleEndian, n)
}

func (p *maplePacketWriter) WriteUInt64(n uint64) (err error) {
	return binary.Write(&p.buf, binary.LittleEndian, n)
}

// WriteString appends a short containing the size of the string and then the string as a byte slice
func (p *maplePacketWriter) WriteString(s string) (n int, err error) {
	err = p.WriteUInt16(uint16(len(s)))
	if err != nil {
		return -1, err
	}

	return p.WriteBytes([]byte(s))
}

// WritePaddedString does not append a short containing the size of the string. Instead, the string is appended, and then /0 padding is added after to fill the length `length`.
func (p *maplePacketWriter) WritePaddedString(s string, length int) (n int, err error) {
	if len(s) > length {
		// Cut the string to fill the fixed length
		s = s[:length]
	}

	strBytes := []byte(s)
	buffer := length - len(strBytes)
	// Append padding
	for i := 0; i < buffer; i++ {
		strBytes = append(strBytes, 0)
	}

	return p.WriteBytes(strBytes)
}
