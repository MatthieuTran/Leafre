package operation

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
