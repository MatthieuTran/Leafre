package reader_test

import (
	"math/rand"
	"testing"

	"github.com/matthieutran/leafre-login/networking/reader"
	"github.com/matthieutran/packet"
)

func TestReadSelect(t *testing.T) {
	var worldId, unknown1, channelId [1]byte
	rand.Read(unknown1[:])
	rand.Read(worldId[:])
	rand.Read(channelId[:])

	p := packet.Packet{}
	p.WriteByte(unknown1[0])
	p.WriteByte(worldId[0])
	p.WriteByte(channelId[0])

	recv := reader.ReadSelectWorld(p)

	// Check if unknown1 matches
	if recv.Unknown1 != unknown1[0] {
		t.Errorf("Expected Unknown1 == %d, actual = %d", unknown1[0], recv.Unknown1)
	}

	// Check if world id matches
	if recv.WorldId != worldId[0] {
		t.Error(worldId, unknown1, channelId, p.Bytes())
		t.Errorf("Expected WorldId == %d, actual = %d", worldId[0], recv.WorldId)
	}

	// Check if channel id matches
	if recv.ChannelId != channelId[0] {
		t.Errorf("Expected ChannelId == %d, actual = %d", channelId[0], recv.ChannelId)
	}
}
