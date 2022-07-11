package writer

import (
	"io"

	"github.com/matthieutran/leafre-login/internal/domain/character"
	"github.com/matthieutran/leafre-login/internal/domain/user"
	"github.com/matthieutran/leafre-login/pkg/packet"
)

// WorldInformation provides information about each available world to the client
var OpCodeSelectWorldResult uint16 = 0xB

type SendSelectWorld struct {
	Result     user.LoginResponse
	Characters character.Characters
}

// WriteSelectWorldResult writes the world user limit information
func WriteSelectWorldResult(w io.Writer, send SendSelectWorld) {
	pw := packet.NewPacketWriter()
	pw.WriteUInt16(OpCodeSelectWorldResult)
	pw.WriteOne(byte(send.Result))

	if send.Result == user.LoginResponseSuccess {
		// Send characters
		pw.WriteOne(byte(len(send.Characters))) // Character count
		for _, c := range send.Characters {
			WriteCharacterStats(pw, c)
			WriteCharacterLook(pw, c)

			pw.WriteOne(0)
			pw.WriteOne(0)
		}

		pw.WriteOne(0)    // SPW
		pw.WriteUInt32(3) // Max number of characters
		pw.WriteUInt32(0)
	}

	// Write world to client
	w.Write(pw.Packet())
}
