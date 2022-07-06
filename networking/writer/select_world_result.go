package writer

import (
	"io"

	"github.com/matthieutran/leafre-login/user"
	"github.com/matthieutran/packet"
)

// WorldInformation provides information about each available world to the client
var OpCodeSelectWorldResult uint16 = 0xB

// WriteSelectWorldResult writes the world user limit information
func WriteSelectWorldResult(w io.Writer, result user.LoginResponse) {
	p := packet.Packet{}
	p.WriteShort(OpCodeSelectWorldResult)
	p.WriteByte(0) // bWarningLevel
	p.WriteByte(0) // bPopulateLevel
	p.WriteByte(byte(result))

	if result == user.LoginResponseSuccess {
		var characters []interface{}
		// Send characters
		p.WriteByte(byte(len(characters))) // Character count
		for range characters {
			// Write stats
			// Write look

			p.WriteByte(0)
			p.WriteByte(0)
		}

		p.WriteByte(0) // SPW
		p.WriteInt(10) // Max number of characters
		p.WriteInt(0)
	}

	// Write world to client
	w.Write(p.Bytes())
}
