package auth

import (
	"github.com/matthieutran/leafre-login/internal/login/net/handler/auth/operation"
	"github.com/matthieutran/packet"
)

type LoginRequest struct {
	username string
	password string
	hwid     []byte
}

func HandleLogin(p packet.Packet) {
	req := readLogin(p)    // Read packet and create LoginRequest struct
	op := req.checkLogin() // Check login

	switch op {
	case operation.Success:
		req.sendSuccess()
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
		username: username,
		password: password,
		hwid:     hwid,
	}
}

// checkLogin validates a login and returns a `LoginRequestCode`
func (s *LoginRequest) checkLogin() operation.LoginRequestCode {
	return operation.Success
}

// sendSuccess announces to the client that login was successful
func (s *LoginRequest) sendSuccess() {
	// response struct

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
