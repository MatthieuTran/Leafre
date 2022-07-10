package writer

import (
	"io"

	"github.com/matthieutran/leafre-login/internal/domain/user"
	"github.com/matthieutran/leafre-login/pkg/packet"
)

type SendCheckLogin struct {
	Result            byte   // LoginResult
	Unknown1          byte   // Unknown1
	Unknown2          uint32 // Unknown2
	ID                uint32 // AccountID
	Name              string // sNexonClubID
	AdminLevel        byte   // AdminLevel
	GMLevel           uint16 // GMLevel
	CountryID         byte   // CountryID
	Gender            byte   // Gender
	PurchaseEXP       byte   // nPurchaseEXP
	ChatUnblockReason byte   // ChatUnblockReason
	ChatUnblockDate   uint64 // dtChatUnblockDate
	RegisterDate      uint64 // dtRegisterDate
	NumOfCharacter    uint32 // nNumOfCharacter
	V44               byte   // v44
	Msg               byte   // sMsg
	SessionKey        uint64 // session key (for preventing remote hacks)
}

var OpCodeCheckPasswordResult uint16 = 0x0

func WriteCheckPasswordResult(w io.Writer, send SendCheckLogin) {
	p := packet.NewPacketWriter()
	p.WriteUInt16(OpCodeCheckPasswordResult)
	p.WriteOne(send.Result)
	p.WriteOne(send.Unknown1)
	p.WriteUInt32(uint32(send.Unknown2))

	if send.Result == byte(user.LoginResponseSuccess) {
		p.WriteUInt32(send.ID)
		p.WriteOne(send.Gender)
		p.WriteOne(send.AdminLevel)
		p.WriteUInt16(send.GMLevel)
		p.WriteOne(send.CountryID)
		p.WriteString(send.Name)
		p.WriteOne(send.PurchaseEXP)
		p.WriteOne(send.ChatUnblockReason)
		p.WriteUInt64(send.ChatUnblockDate)
		p.WriteUInt64(send.RegisterDate)
		p.WriteUInt32(send.NumOfCharacter)
		p.WriteOne(send.V44)
		p.WriteOne(send.Msg)
		p.WriteUInt64(send.SessionKey)
	}

	w.Write(p.Packet())
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
