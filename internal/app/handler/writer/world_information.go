package writer

import (
	"io"

	"github.com/matthieutran/leafre-login/internal/domain/channel"
	"github.com/matthieutran/leafre-login/internal/domain/world"
	"github.com/matthieutran/leafre-login/pkg/packet"
)

// OpCodeWorldInformation provides information about each available world to the client
var OpCodeWorldInformation uint16 = 0xA

// OpCodeLatestConnectedWorld provides the client with its last seleted world
var OpCodeLatestConnectedWorld uint16 = 0x18

// WriteWorldInformation writes information about a *single* world
func WriteWorldInformation(w io.Writer, world world.World, channels channel.Channels) {
	p := packet.NewPacketWriter()
	p.WriteUInt16(OpCodeWorldInformation)
	p.WriteOne(world.ID)
	p.WriteString(world.Name)
	p.WriteOne(world.State)
	p.WriteString(world.EventDesc) // WorldEventDesc
	p.WriteUInt16(world.EventEXP)  // WorldEventEXP_WSE, WorldSpecificEvent
	p.WriteUInt16(world.EventDrop) // WorldEventDrop_WSE, WorldSpecificEvent
	p.WriteOne(world.BlockCharCreation)

	// Number of channels
	p.WriteOne(byte(len(channels)))

	// Write channel information for world
	for _, channel := range channels {
		p.WriteString(channel.ID)
		p.WriteUInt32(uint32(channel.UserNo))
		p.WriteOne(channel.WorldID)
		p.WriteOne(channel.ChannelID)
		p.WriteOne(channel.AdultChannel)
	}

	p.WriteUInt16(world.Balloon) // TODO: Balloon

	// Write world to client
	w.Write(p.Packet())
}

// WriteWorldInformationDone writes the signal that the world information sending is done
func WriteWorldInformationDone(w io.Writer) {
	p := packet.NewPacketWriter()
	p.WriteUInt16(OpCodeWorldInformation)
	p.WriteOne(0xFF)

	w.Write(p.Packet())
}

type SendLatestConnectedWorld struct {
	LatestConnectedWorld uint32
}

// WriteLatestConnectedWorld writes the user's latest connected world ID
func WriteLatestConnectedWorld(w io.Writer, send SendLatestConnectedWorld) {
	p := packet.NewPacketWriter()
	p.WriteUInt16(OpCodeLatestConnectedWorld)
	p.WriteUInt32(send.LatestConnectedWorld)

	w.Write(p.Packet())
}
