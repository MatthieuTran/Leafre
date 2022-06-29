package operation

type LoginRequestCode byte

const (
	Success                            LoginRequestCode = 0x0
	TempBlocked                        LoginRequestCode = 0x1
	Blocked                            LoginRequestCode = 0x2
	Abandoned                          LoginRequestCode = 0x3
	IncorrectPassword                  LoginRequestCode = 0x4
	NotRegistered                      LoginRequestCode = 0x5
	DBFail                             LoginRequestCode = 0x6
	AlreadyConnected                   LoginRequestCode = 0x7
	NotConnectableWorld                LoginRequestCode = 0x8
	Unknown                            LoginRequestCode = 0x9
	Timeout                            LoginRequestCode = 0xA
	NotAdult                           LoginRequestCode = 0xB
	AuthFail                           LoginRequestCode = 0xC
	ImpossibleIP                       LoginRequestCode = 0xD
	NotAuthorizedNexonID               LoginRequestCode = 0xE
	NoNexonID                          LoginRequestCode = 0xF
	NotAuthorized                      LoginRequestCode = 0x10
	InvalidRegionInfo                  LoginRequestCode = 0x11
	InvalidBirthDate                   LoginRequestCode = 0x12
	PassportSuspended                  LoginRequestCode = 0x13
	IncorrectSSN2                      LoginRequestCode = 0x14
	WebAuthNeeded                      LoginRequestCode = 0x15
	DeleteCharacterFailedOnGuildMaster LoginRequestCode = 0x16
	NotagreedEULA                      LoginRequestCode = 0x17
	DeleteCharacterFailedEngaged       LoginRequestCode = 0x18
	IncorrectSPW                       LoginRequestCode = 0x14
	SamePasswordAndSPW                 LoginRequestCode = 0x16
	SamePincodeAndSPW                  LoginRequestCode = 0x17
	RegisterLimitedIP                  LoginRequestCode = 0x19
	RequestedCharacterTransfer         LoginRequestCode = 0x1A
	CashUserCannotUseSimpleClient      LoginRequestCode = 0x1B
	DeleteCharacterFailedOnFamily      LoginRequestCode = 0x1D
	InvalidCharacterName               LoginRequestCode = 0x1E
	IncorrectSSN                       LoginRequestCode = 0x1F
	SSNConfirmFailed                   LoginRequestCode = 0x20
	SSNNotConfirmed                    LoginRequestCode = 0x21
	WorldTooBusy                       LoginRequestCode = 0x22
	OTPReissuing                       LoginRequestCode = 0x23
	OTPInfoNotExist                    LoginRequestCode = 0x24
)
