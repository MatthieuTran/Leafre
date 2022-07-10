package reader

import "github.com/matthieutran/leafre-login/pkg/packet"

// RecvCheckLogin defines the structure returned by the client including the client's username and password
type RecvCheckLogin struct {
	Username       string
	Password       string
	MachineId      []byte
	GameRoomClient uint32
	GameStartMode  byte
	Unknown1       byte
	Unknown2       byte
	PartnerCode    uint32
}

// ReadLogin parses the packet into a processable struct
func ReadLogin(p packet.Packet) (recv RecvCheckLogin) {
	pr := packet.NewPacketReader(p)
	recv.Username = pr.ReadString()       // Username
	recv.Password = pr.ReadString()       // Password
	recv.MachineId = pr.ReadBytes(16)     // Machine ID
	recv.GameRoomClient = pr.ReadUInt32() // Game Room Client
	recv.GameStartMode = pr.ReadOne()     // GameStartMode
	recv.Unknown1 = pr.ReadOne()          // Unknown1
	recv.Unknown2 = pr.ReadOne()          // Unknown2
	recv.PartnerCode = pr.ReadUInt32()    // PartnerCode

	return
}
