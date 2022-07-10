package reader

import "github.com/matthieutran/leafre-login/pkg/packet"

type RecvCheckUserLimit struct {
	WorldId  byte
	Unknown1 byte
}

func ReadCheckUserLimit(p packet.Packet) (recv RecvCheckUserLimit) {
	pr := packet.NewPacketReader(p)
	recv.WorldId = pr.ReadOne()
	recv.Unknown1 = pr.ReadOne()

	return
}
