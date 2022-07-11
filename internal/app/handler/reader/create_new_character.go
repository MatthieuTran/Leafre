package reader

import (
	"github.com/matthieutran/leafre-login/pkg/packet"
)

type RecvCreateNewCharacter struct {
	Name      string
	Race      uint32
	SubJob    uint16
	Face      uint32
	Hair      uint32
	HairColor uint32
	Skin      uint32
	Coat      uint32
	Pants     uint32
	Shoes     uint32
	Weapon    uint32
	Gender    byte
}

func ReadCreateNewCharacter(p packet.Packet) (recv RecvCreateNewCharacter) {
	pr := packet.NewPacketReader(p)
	recv.Name = pr.ReadString()
	recv.Race = pr.ReadUInt32()
	recv.SubJob = pr.ReadUInt16()
	recv.Face = pr.ReadUInt32()
	recv.Hair = pr.ReadUInt32()
	recv.HairColor = pr.ReadUInt32()
	recv.Skin = pr.ReadUInt32()
	recv.Coat = pr.ReadUInt32()
	recv.Pants = pr.ReadUInt32()
	recv.Shoes = pr.ReadUInt32()
	recv.Weapon = pr.ReadUInt32()
	recv.Gender = pr.ReadOne()

	return
}
