package writer

import (
	"io"

	login "github.com/matthieutran/leafre-login"
	"github.com/matthieutran/packet"
)

// WorldInformation provides information about each available world to the client
var OpCodeWorldInformation uint16 = 0xA

// WriteWorldInformation writes information about a *single* world
func WriteWorldInformation(w io.Writer, world login.World, channels login.Channels) {
	p := packet.Packet{}
	p.WriteShort(OpCodeWorldInformation)
	p.WriteByte(world.Id)
	p.WriteString(world.Name)
	p.WriteByte(world.State)
	p.WriteString("") // WorldEventDesc
	p.WriteShort(0)   // WorldEventEXP_WSE, WorldSpecificEvent
	p.WriteShort(0)   // WorldEventDrop_WSE, WorldSpecific Event
	p.WriteByte(world.BlockCharCreation)

	// Number of channels
	p.WriteByte(byte(len(channels)))

	// Write channel information for world
	for _, channel := range channels {
		p.WriteString(channel.Id)
		p.WriteInt(uint32(channel.UserNo))
		p.WriteByte(channel.WorldId)
		p.WriteByte(channel.ChannelId)
		p.WriteByte(channel.AdultChannel)
	}

	p.WriteShort(0) // TODO: Balloon

	// Write world to client
	w.Write(p.Bytes())
}

// WriteWorldInformationDone writes the signal that the world information sending is done
func WriteWorldInformationDone(w io.Writer) {
	p := packet.Packet{}
	p.WriteShort(OpCodeWorldInformation)
	p.WriteByte(0xFF)

	w.Write(p.Bytes())
}

var OpCodeLatestConnectedWorld uint16 = 0x18

// WriteLatestConnectedWorld writes the user's latest connected world ID
func WriteLatestConnectedWorld(w io.Writer, latestConnectedWorld int) {
	p := packet.Packet{}
	p.WriteShort(OpCodeLatestConnectedWorld)
	p.WriteInt(uint32(latestConnectedWorld))

	w.Write(p.Bytes())
}
