package user

type LoginForm struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	MachineId string `json:"machine_id"`
}

// Service provides access to a user store
type Service interface {
	// Login validates the login details in the `UserForm` object and returns the user's object and error (where applicable)
	Login(LoginForm) (user User, code LoginResponse)

	// GetById fetches a user by its ID
	GetById(int) (user User, err error)
}

type LoginResponse uint16

const (
	LoginResponseSuccess                            LoginResponse = 0x0
	LoginResponseTempBlocked                        LoginResponse = 0x1
	LoginResponseBlocked                            LoginResponse = 0x2
	LoginResponseAbandoned                          LoginResponse = 0x3
	LoginResponseIncorrectPassword                  LoginResponse = 0x4
	LoginResponseNotRegistered                      LoginResponse = 0x5
	LoginResponseDBFail                             LoginResponse = 0x6
	LoginResponseAlreadyConnected                   LoginResponse = 0x7
	LoginResponseNotConnectableWorld                LoginResponse = 0x8
	LoginResponseUnknown                            LoginResponse = 0x9
	LoginResponseTimeout                            LoginResponse = 0xA
	LoginResponseNotAdult                           LoginResponse = 0xB
	LoginResponseAuthFail                           LoginResponse = 0xC
	LoginResponseImpossibleIP                       LoginResponse = 0xD
	LoginResponseNotAuthorizedNexonID               LoginResponse = 0xE
	LoginResponseNoNexonID                          LoginResponse = 0xF
	LoginResponseNotAuthorized                      LoginResponse = 0x10
	LoginResponseInvalidRegionInfo                  LoginResponse = 0x11
	LoginResponseInvalidBirthDate                   LoginResponse = 0x12
	LoginResponsePassportSuspended                  LoginResponse = 0x13
	LoginResponseIncorrectSSN2                      LoginResponse = 0x14
	LoginResponseWebAuthNeeded                      LoginResponse = 0x15
	LoginResponseDeleteCharacterFailedOnGuildMaster LoginResponse = 0x16
	LoginResponseNotagreedEULA                      LoginResponse = 0x17
	LoginResponseDeleteCharacterFailedEngaged       LoginResponse = 0x18
	LoginResponseIncorrectSPW                       LoginResponse = 0x14
	LoginResponseSamePasswordAndSPW                 LoginResponse = 0x16
	LoginResponseSamePincodeAndSPW                  LoginResponse = 0x17
	LoginResponseRegisterLimitedIP                  LoginResponse = 0x19
	LoginResponseRequestedCharacterTransfer         LoginResponse = 0x1A
	LoginResponseCashUserCannotUseSimpleClient      LoginResponse = 0x1B
	LoginResponseDeleteCharacterFailedOnFamily      LoginResponse = 0x1D
	LoginResponseInvalidCharacterName               LoginResponse = 0x1E
	LoginResponseIncorrectSSN                       LoginResponse = 0x1F
	LoginResponseSSNConfirmFailed                   LoginResponse = 0x20
	LoginResponseSSNNotConfirmed                    LoginResponse = 0x21
	LoginResponseWorldTooBusy                       LoginResponse = 0x22
	LoginResponseOTPReissuing                       LoginResponse = 0x23
	LoginResponseOTPInfoNotExist                    LoginResponse = 0x24
)
