package reader

import (
	"github.com/matthieutran/packet"
)

type RecvSelectWorld struct {
	Unknown1  byte `json:"unknown1"`
	WorldId   byte `json:"world_id"`
	ChannelId byte `json:"channel_id"`
}

func ReadSelectWorld(p packet.Packet) (res RecvSelectWorld) {
	res.Unknown1 = p.ReadBytes(1)[0]
	res.WorldId = p.ReadBytes(1)[0]
	res.ChannelId = p.ReadBytes(1)[0]

	return
}
