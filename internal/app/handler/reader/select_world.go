package reader

import "github.com/matthieutran/leafre-login/pkg/packet"

type RecvSelectWorld struct {
	Unknown1  byte `json:"unknown1"`
	WorldID   byte `json:"world_id"`
	ChannelId byte `json:"channel_id"`
}

func ReadSelectWorld(p packet.Packet) (recv RecvSelectWorld) {
	pr := packet.NewPacketReader(p)
	recv.Unknown1 = pr.ReadOne()
	recv.WorldID = pr.ReadOne()
	recv.ChannelId = pr.ReadOne()

	return
}
