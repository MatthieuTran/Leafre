package packet_test

import (
	"bytes"
	"math/rand"
	"testing"

	"github.com/matthieutran/leafre-login/pkg/packet"
)

func TestPacketReader(t *testing.T) {
	var buf [24]byte
	rand.Read(buf[:])

	p := packet.NewPacketReader(buf[:])

	// Reading a UInt16 should return the first two bytes of the buffer
	expectedShort := uint16(buf[0]) | uint16(buf[1])<<8
	actualShort := p.ReadUInt16()
	if actualShort != expectedShort {
		t.Errorf("Expected buf[0:2] == %d, actual = %d", expectedShort, actualShort)
	}

	// Reading a UInt32 should return the next four bytes of the buffer
	expectedInt := uint32(buf[2]) |
		uint32(buf[3])<<8 |
		uint32(buf[4])<<16 |
		uint32(buf[5])<<24
	actualInt := p.ReadUInt32()
	if actualInt != expectedInt {
		t.Errorf("Expected buf[2:6] == %d, actual = %d", expectedInt, actualInt)
	}

	// Reading a UInt64 should return the next eight bytes of the buffer
	expectedLong := uint64(buf[6]) |
		uint64(buf[7])<<8 |
		uint64(buf[8])<<16 |
		uint64(buf[9])<<24 |
		uint64(buf[10])<<32 |
		uint64(buf[11])<<40 |
		uint64(buf[12])<<48 |
		uint64(buf[13])<<56
	actualLong := p.ReadUInt64()
	if actualLong != expectedLong {
		t.Errorf("Expected buf[7:14] == %d, actual = %d", expectedLong, actualLong)
	}

	// Reading a byte should return the next byte of the buffer
	expectedByte := buf[14]
	actualByte := p.ReadOne()
	if actualByte != expectedByte {
		t.Errorf("Expected buf[14] == %d, actual = %d", expectedByte, actualByte)
	}

	// Reading nine bytes should return the next nine bytes of the buffer
	expectedBytes := buf[15:24]
	actualBytes := p.ReadBytes(9)
	if !bytes.Equal(actualBytes[:], expectedBytes) {
		t.Errorf("Expected buf[15:24] == %d, actual = %d", expectedBytes, actualBytes)
	}
}
