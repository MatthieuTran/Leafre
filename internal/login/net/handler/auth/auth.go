package auth

import (
	"log"

	"github.com/matthieutran/duey"
	"github.com/matthieutran/leafre-login/pkg/operation"
	"github.com/matthieutran/packet"
	"github.com/matthieutran/tcpserve"
)

type LoginRequest struct {
	Username string
	Password string
	Hwid     []byte
}

func HandleLogin(s *tcpserve.Session, es *duey.EventStreamer, p packet.Packet) {
	req := readLogin(p)        // Read packet and create LoginRequest struct
	res := checkLogin(es, req) // Check login
	log.Println(res)

	switch res.Code {
	case operation.Success:
		sendSuccess(s, res, req.Username)
	default:
		sendFailed(s, res)
	}
}

// readLogin parses the packet into a processable struct
func readLogin(p packet.Packet) *LoginRequest {
	_, username := p.ReadString() // Username
	_, password := p.ReadString() // Password
	hwid := p.ReadBytes(16)       // Machine ID
	_ = p.ReadInt()               // Game Room Client
	_ = p.ReadBytes(1)            // GameStartMode
	_ = p.ReadBytes(1)            // Unknown1
	_ = p.ReadBytes(1)            // Unknown2
	_ = p.ReadInt()               // PartnerCode

	return &LoginRequest{
		Username: username,
		Password: password,
		Hwid:     hwid,
	}
}

func sendFailed(s *tcpserve.Session, res loginResponse) {
	p := packet.Packet{}
	p.WriteShort(0)             // Header
	p.WriteByte(byte(res.Code)) // Result
	p.WriteByte(0)              // Unknown1
	p.WriteInt(0)               // Unknown2
	s.Write(p.Bytes())
}

// sendSuccess announces to the client that login was successful
func sendSuccess(s *tcpserve.Session, res loginResponse, username string) {
	p := packet.Packet{}
	p.WriteShort(0x00)
	p.WriteByte(byte(res.Code)) // Result
	p.WriteByte(0)              // Unknown1
	p.WriteInt(0)               // Unknown2
	p.WriteInt(uint32(res.Id))  // AccountID
	p.WriteByte(0)              // Gender TODO: change me
	p.WriteByte(0)              // AdminLevel
	p.WriteShort(0)             // GM Level
	p.WriteByte(0)              // nCountryID
	p.WriteString(username)     // sNexonClubID
	p.WriteByte(0)              // nPurchaseEXP
	p.WriteByte(0)              // ChatUnblockReason
	p.WriteLong(0)              // dtChatUnblockDate
	p.WriteLong(0)              // dtRegisterDate
	p.WriteInt(4)               // nNumOfCharacter
	p.WriteByte(1)              // v44
	p.WriteByte(0)              // sMsg
	p.WriteLong(0)              // session key (for preventing remote hacks)
	s.Write(p.Bytes())

	/**
	byte - result
	byte - unk1
	byte - unk2

	if result == 0
		int - acc id
		byte - gender (1 male, 0 female)

		byte - grade code bitflag (admin level)
			AdminLevel1 = 0x1,
			AdminLevel2 = 0x2,
			AdminLevel3 = 0x4,
			AdminLevel4 = 0x8,
			AdminLevel5 = 0x10,
			AdminLevel10 = 0x20

		byte - sub grade code bitflag (gm level)
			PrimaryTrace = 0x1,
			SecondaryTrace = 0x2,
			AdminClient = 0x4,
			MobMoveObserve = 0x8,
			ManagerAccount = 0x10,
			OutSourceSuperGM = 0x20,
			OutSourceGM = 0x40,
			UserGM = 0x80,
			TesterAccount = 0x100

		byte - country id
		string - username (nexon club id)
		byte
		byte
		long
		long
		int - num of character (not sure)
		byte - set as 1 (not sure)
		byte

		long - session key (for preventing remote hacks)
	end

	**/
}
