package reader

import (
	"github.com/matthieutran/packet"
)

// RecvCheckLogin defines the structure returned by the client including the client's username and password
type RecvCheckLogin struct {
	Username       string
	Password       string
	MachineId      []byte
	GameRoomClient int
	GameStartMode  byte
	Unknown1       byte
	Unknown2       byte
	PartnerCode    int
}

// ReadLogin parses the packet into a processable struct
func ReadLogin(p packet.Packet) *RecvCheckLogin {
	_, username := p.ReadString() // Username
	_, password := p.ReadString() // Password
	machineId := p.ReadBytes(16)  // Machine ID
	grm := p.ReadInt()            // Game Room Client
	gsm := p.ReadBytes(1)         // GameStartMode
	unk1 := p.ReadBytes(1)        // Unknown1
	unk2 := p.ReadBytes(1)        // Unknown2
	pc := p.ReadInt()             // PartnerCode

	return &RecvCheckLogin{
		Username:       username,
		Password:       password,
		MachineId:      machineId,
		GameRoomClient: int(grm),
		GameStartMode:  gsm[0],
		Unknown1:       unk1[0],
		Unknown2:       unk2[0],
		PartnerCode:    int(pc),
	}
}
