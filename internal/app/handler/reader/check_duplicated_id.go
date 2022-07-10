package reader

import "github.com/matthieutran/leafre-login/pkg/packet"

type RecvCheckDuplicatedID struct {
	Name string
}

// ReadLogin parses the packet into a processable struct
func ReadCheckDuplicatedID(p packet.Packet) (recv RecvCheckDuplicatedID) {
	pr := packet.NewPacketReader(p)
	recv.Name = pr.ReadString()

	return
}
