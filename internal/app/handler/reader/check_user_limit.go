package reader

import "github.com/matthieutran/leafre-login/pkg/packet"

type RecvCheckUserLimit struct {
	WorldId  byte
	Unknown1 byte
}

func ReadCheckUserLimit(p packet.Packet) (res RecvCheckUserLimit) {
	pr := packet.NewPacketReader(p)
	res.WorldId = pr.ReadOne()
	res.Unknown1 = pr.ReadOne()

	return
}
