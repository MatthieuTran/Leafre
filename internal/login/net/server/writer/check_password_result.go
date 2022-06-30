package writer

import (
	"io"

	"github.com/matthieutran/packet"
)

var OpCodeCheckPasswordResult uint16 = 0x0

type CodeLoginRequest uint16

const (
	LoginRequestSuccess                            CodeLoginRequest = 0x0
	LoginRequestTempBlocked                        CodeLoginRequest = 0x1
	LoginRequestBlocked                            CodeLoginRequest = 0x2
	LoginRequestAbandoned                          CodeLoginRequest = 0x3
	LoginRequestIncorrectPassword                  CodeLoginRequest = 0x4
	LoginRequestNotRegistered                      CodeLoginRequest = 0x5
	LoginRequestDBFail                             CodeLoginRequest = 0x6
	LoginRequestAlreadyConnected                   CodeLoginRequest = 0x7
	LoginRequestNotConnectableWorld                CodeLoginRequest = 0x8
	LoginRequestUnknown                            CodeLoginRequest = 0x9
	LoginRequestTimeout                            CodeLoginRequest = 0xA
	LoginRequestNotAdult                           CodeLoginRequest = 0xB
	LoginRequestAuthFail                           CodeLoginRequest = 0xC
	LoginRequestImpossibleIP                       CodeLoginRequest = 0xD
	LoginRequestNotAuthorizedNexonID               CodeLoginRequest = 0xE
	LoginRequestNoNexonID                          CodeLoginRequest = 0xF
	LoginRequestNotAuthorized                      CodeLoginRequest = 0x10
	LoginRequestInvalidRegionInfo                  CodeLoginRequest = 0x11
	LoginRequestInvalidBirthDate                   CodeLoginRequest = 0x12
	LoginRequestPassportSuspended                  CodeLoginRequest = 0x13
	LoginRequestIncorrectSSN2                      CodeLoginRequest = 0x14
	LoginRequestWebAuthNeeded                      CodeLoginRequest = 0x15
	LoginRequestDeleteCharacterFailedOnGuildMaster CodeLoginRequest = 0x16
	LoginRequestNotagreedEULA                      CodeLoginRequest = 0x17
	LoginRequestDeleteCharacterFailedEngaged       CodeLoginRequest = 0x18
	LoginRequestIncorrectSPW                       CodeLoginRequest = 0x14
	LoginRequestSamePasswordAndSPW                 CodeLoginRequest = 0x16
	LoginRequestSamePincodeAndSPW                  CodeLoginRequest = 0x17
	LoginRequestRegisterLimitedIP                  CodeLoginRequest = 0x19
	LoginRequestRequestedCharacterTransfer         CodeLoginRequest = 0x1A
	LoginRequestCashUserCannotUseSimpleClient      CodeLoginRequest = 0x1B
	LoginRequestDeleteCharacterFailedOnFamily      CodeLoginRequest = 0x1D
	LoginRequestInvalidCharacterName               CodeLoginRequest = 0x1E
	LoginRequestIncorrectSSN                       CodeLoginRequest = 0x1F
	LoginRequestSSNConfirmFailed                   CodeLoginRequest = 0x20
	LoginRequestSSNNotConfirmed                    CodeLoginRequest = 0x21
	LoginRequestWorldTooBusy                       CodeLoginRequest = 0x22
	LoginRequestOTPReissuing                       CodeLoginRequest = 0x23
	LoginRequestOTPInfoNotExist                    CodeLoginRequest = 0x24
)

func WriteCheckPasswordResult(s io.Writer, res CodeLoginRequest, id int, username string) {
	p := packet.Packet{}
	p.WriteShort(OpCodeCheckPasswordResult) // Header
	p.WriteByte(byte(res))                  // Result
	p.WriteByte(0)                          // Unknown1
	p.WriteInt(0)                           // Unknown2

	if res == LoginRequestSuccess {
		p.WriteInt(uint32(id))  // AccountID
		p.WriteByte(0)          // Gender TODO: change me
		p.WriteByte(0)          // AdminLevel
		p.WriteShort(0)         // GM Level
		p.WriteByte(0)          // nCountryID
		p.WriteString(username) // sNexonClubID
		p.WriteByte(0)          // nPurchaseEXP
		p.WriteByte(0)          // ChatUnblockReason
		p.WriteLong(0)          // dtChatUnblockDate
		p.WriteLong(0)          // dtRegisterDate
		p.WriteInt(4)           // nNumOfCharacter
		p.WriteByte(1)          // v44
		p.WriteByte(0)          // sMsg
		p.WriteLong(0)          // session key (for preventing remote hacks)
	}

	s.Write(p.Bytes())
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
