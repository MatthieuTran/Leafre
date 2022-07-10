package reader_test

import (
	"math/rand"
	"testing"

	"github.com/matthieutran/leafre-login/internal/app/handler/reader"
	"github.com/matthieutran/leafre-login/pkg/packet"
)

func TestReadSelect(t *testing.T) {
	var worldId, unknown1, channelId [1]byte
	rand.Read(unknown1[:])
	rand.Read(worldId[:])
	rand.Read(channelId[:])

	pw := packet.NewPacketWriter()
	pw.WriteOne(unknown1[0])
	pw.WriteOne(worldId[0])
	pw.WriteOne(channelId[0])

	recv := reader.ReadSelectWorld(pw.Packet())

	// Check if unknown1 matches
	if recv.Unknown1 != unknown1[0] {
		t.Errorf("Expected Unknown1 == %d, actual = %d", unknown1[0], recv.Unknown1)
	}

	// Check if world id matches
	if recv.WorldID != worldId[0] {
		t.Error(worldId, unknown1, channelId, pw.Packet())
		t.Errorf("Expected WorldId == %d, actual = %d", worldId[0], recv.WorldID)
	}

	// Check if channel id matches
	if recv.ChannelId != channelId[0] {
		t.Errorf("Expected ChannelId == %d, actual = %d", channelId[0], recv.ChannelId)
	}
}
