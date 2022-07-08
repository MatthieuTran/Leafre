package packet_test

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"testing"

	"github.com/matthieutran/leafre-login/pkg/packet"
)

// TestPacketWriter tests the PacketWriter and ensures what is put in is what is received
func TestPacketWriter(t *testing.T) {
	var buf [24]byte
	rand.Read(buf[:])

	p := packet.NewPacketWriter()

	indx := 0
	// WriteUInt16
	expectedShort := uint16(buf[indx]) | uint16(buf[indx+1])<<8
	p.WriteUInt16(expectedShort)
	actualShort := p.Packet()[indx : indx+2]
	if binary.LittleEndian.Uint16(actualShort) != expectedShort {
		t.Errorf("Expected buf[%d:%d] == %d, actual = %d", indx, indx+2, expectedShort, binary.LittleEndian.Uint16(actualShort))
	}
	indx += 2

	// WriteUInt32
	expectedInt := uint32(buf[indx]) |
		uint32(buf[indx+1])<<8 |
		uint32(buf[indx+2])<<16 |
		uint32(buf[indx+3])<<24
	p.WriteUInt32(expectedInt)
	actualInt := p.Packet()[indx : indx+4]
	if binary.LittleEndian.Uint32(actualInt) != expectedInt {
		t.Errorf("Expected buf[%d:%d] == %d, actual = %d", indx, indx+4, expectedInt, binary.LittleEndian.Uint32(actualInt))
	}
	indx += 4

	// WriteUInt64
	expectedLong := uint64(buf[indx]) |
		uint64(buf[indx+1])<<8 |
		uint64(buf[indx+2])<<16 |
		uint64(buf[indx+3])<<24 |
		uint64(buf[indx+4])<<32 |
		uint64(buf[indx+5])<<40 |
		uint64(buf[indx+6])<<48 |
		uint64(buf[indx+7])<<56
	p.WriteUInt64(expectedLong)
	actualLong := p.Packet()[indx : indx+8]
	if binary.LittleEndian.Uint64(actualLong) != expectedLong {
		t.Errorf("Expected buf[%d:%d] == %d, actual = %d", indx, indx+8, expectedLong, actualLong)
	}
	indx += 8

	// WriteByte
	p.WriteByte(buf[indx])
	expectedByte := buf[indx]
	actualByte := p.Packet()[indx]
	if actualByte != expectedByte {
		t.Errorf("Expected buf[%d] == %d, actual = %d", indx, expectedByte, actualByte)
	}
	indx += 1

	// Write
	expectedBytes := buf[indx : indx+9]
	p.Write(expectedBytes[:])
	actualBytes := p.Packet()[indx : indx+9]
	if !bytes.Equal(actualBytes[:], expectedBytes) {
		t.Errorf("Expected buf[15:24] == %d, actual = %d", expectedBytes, actualBytes)
	}
	indx += 9
}
