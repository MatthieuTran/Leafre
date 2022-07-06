package writer

import (
	"io"

	"github.com/matthieutran/leafre-login/user"
	"github.com/matthieutran/packet"
)

var OpCodeCheckPasswordResult uint16 = 0x0

func WriteCheckPasswordResult(w io.Writer, res user.LoginResponse, u user.User) {
	p := packet.Packet{}
	p.WriteShort(OpCodeCheckPasswordResult) // Header
	p.WriteByte(byte(res))                  // Result
	p.WriteByte(0)                          // Unknown1
	p.WriteInt(0)                           // Unknown2

	if res == user.LoginResponseSuccess {
		p.WriteInt(uint32(u.Id))  // AccountID
		p.WriteByte(u.Gender)     // Gender TODO: change me
		p.WriteByte(0)            // AdminLevel
		p.WriteShort(0)           // GM Level
		p.WriteByte(0)            // nCountryID
		p.WriteString(u.Username) // sNexonClubID
		p.WriteByte(0)            // nPurchaseEXP
		p.WriteByte(0)            // ChatUnblockReason
		p.WriteLong(0)            // dtChatUnblockDate
		p.WriteLong(0)            // dtRegisterDate
		p.WriteInt(4)             // nNumOfCharacter
		p.WriteByte(1)            // v44
		p.WriteByte(0)            // sMsg
		p.WriteLong(0)            // session key (for preventing remote hacks)
	}

	w.Write(p.Bytes())
}

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
