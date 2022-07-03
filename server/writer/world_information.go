package writer

import (
	"io"

	login "github.com/matthieutran/leafre-login"
)

// WorldInformation provides information about each available world to the client
var OpCodeWorldInformation uint16 = 0xA

func WriteWorldInformation(w io.Writer, world login.World) {

}
