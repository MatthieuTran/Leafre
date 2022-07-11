package reader_test

import (
	"math/rand"
	"testing"

	"github.com/matthieutran/leafre-login/internal/app/handler/reader"
	"github.com/matthieutran/leafre-login/pkg/packet"
)

func TestReadCheckUserLimit(t *testing.T) {
	var worldId, unknown1 [1]byte
	rand.Read(worldId[:])
	rand.Read(unknown1[:])

	pw := packet.NewPacketWriter()
	pw.WriteOne(worldId[0])
	pw.WriteOne(unknown1[0])

	recv := reader.ReadCheckUserLimit(pw.Packet())

	// Check if world id matches
	if recv.WorldId != worldId[0] {
		t.Errorf("Expected WorldId == %d, actual = %d", worldId, recv.WorldId)
	}

	// Check if unknown1 matches
	if recv.Unknown1 != unknown1[0] {
		t.Errorf("Expected Unknown1 == %d, actual = %d", unknown1[0], recv.Unknown1)
	}
}

// TestReadCheckUserLimitInt tests that if we have a packet with an int instead of two bytes,
// we should be receiving the first two bytes of the int instead of the entire int
func TestReadCheckUserLimitInt(t *testing.T) {
	// Use a size-4 byte slice to simulate an integer
	var num [4]byte
	rand.Read(num[:])

	pw := packet.NewPacketWriter()
	pw.WriteBytes(num[:])

	recv := reader.ReadCheckUserLimit(pw.Packet())

	// WorldId should be the first byte in the integer
	if recv.WorldId != num[0] {
		t.Errorf("Expected WorldId == %d, actual = %d", num[0], recv.WorldId)
	}

	// Unknown1 should be the second byte in the integer
	if recv.Unknown1 != num[1] {
		t.Errorf("Expected Unknown1 == %d, actual = %d", num[1], recv.Unknown1)
	}
}
