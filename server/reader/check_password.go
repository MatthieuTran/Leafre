package reader

import (
	"github.com/matthieutran/packet"
)

// RecvCheckLogin defines the structure returned by the client including the client's username and password
type RecvCheckLogin struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	MachineId      []byte `json:"machine_id"`
	GameRoomClient int    `json:"game_room_client"`
	GameStartMode  byte   `json:"game_start_mode"`
	Unknown1       byte   `json:"unknown1"`
	Unknown2       byte   `json:"unknown2"`
	PartnerCode    int    `json:"partner_code"`
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
