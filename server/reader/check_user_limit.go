package reader

import "github.com/matthieutran/packet"

type RecvCheckUserLimit struct {
	WorldId  byte `json:"world_id"`
	Unknown1 byte `json:"unknown1"`
}

func ReadCheckUserLimit(p packet.Packet) (res RecvCheckUserLimit) {
	res.WorldId = p.ReadBytes(1)[0]
	res.Unknown1 = p.ReadBytes(1)[0]

	return
}
