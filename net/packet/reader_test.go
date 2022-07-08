package packet_test

import (
	"bytes"
	"math/rand"
	"testing"

	"github.com/matthieutran/leafre-login/net/packet"
)

func TestPacketReader(t *testing.T) {
	var buf [24]byte
	rand.Read(buf[:])

	p := packet.NewPacketReader(buf[:])

	// Reading a UInt16 should return the first two bytes of the buffer
	var actualShort uint16
	expectedShort := uint16(buf[0]) | uint16(buf[1])<<8
	p.ReadUInt16(&actualShort)
	if actualShort != expectedShort {
		t.Errorf("Expected buf[0:2] == %d, actual = %d", expectedShort, actualShort)
	}

	// Reading a UInt32 should return the next four bytes of the buffer
	var actualInt uint32
	expectedInt := uint32(buf[2]) |
		uint32(buf[3])<<8 |
		uint32(buf[4])<<16 |
		uint32(buf[5])<<24
	p.ReadUInt32(&actualInt)
	if actualInt != expectedInt {
		t.Errorf("Expected buf[2:6] == %d, actual = %d", expectedInt, actualInt)
	}

	// Reading a UInt64 should return the next eight bytes of the buffer
	var actualLong uint64
	expectedLong := uint64(buf[6]) |
		uint64(buf[7])<<8 |
		uint64(buf[8])<<16 |
		uint64(buf[9])<<24 |
		uint64(buf[10])<<32 |
		uint64(buf[11])<<40 |
		uint64(buf[12])<<48 |
		uint64(buf[13])<<56
	p.ReadUInt64(&actualLong)
	if actualLong != expectedLong {
		t.Errorf("Expected buf[7:14] == %d, actual = %d", expectedLong, actualLong)
	}

	// Reading a byte should return the next byte of the buffer
	actualByte, _ := p.ReadByte()
	expectedByte := buf[14]
	if actualByte != expectedByte {
		t.Errorf("Expected buf[14] == %d, actual = %d", expectedByte, actualByte)
	}

	// Reading nine bytes should return the next nine bytes of the buffer
	var actualBytes [9]byte
	p.Read(actualBytes[:])
	expectedBytes := buf[15:24]
	if !bytes.Equal(actualBytes[:], expectedBytes) {
		t.Errorf("Expected buf[15:24] == %d, actual = %d", expectedBytes, actualBytes)
	}

	// Reading a string should return correctly
	var strLength [1]byte
	rand.Read(strLength[:])
	str := randString(int(strLength[0]))
	strBytes := []byte(str)
	strBuffer := []byte{strLength[0]}
	strBuffer = append(strBuffer, strBytes...)

	p = packet.NewPacketReader(strBuffer)
	actualStr, _ := p.ReadString()
	if actualStr != str {
		t.Errorf("Expected pStr == %s, actual = %s", str, actualStr)
	}
}

func randString(n int) string {
	var parts = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	res := make([]rune, n)
	for i := range res {
		res[i] = parts[rand.Intn(len(parts))]
	}

	return string(res)
}
