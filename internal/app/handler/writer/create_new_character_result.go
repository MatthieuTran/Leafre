package writer

import (
	"io"

	"github.com/matthieutran/leafre-login/internal/domain/character"
	"github.com/matthieutran/leafre-login/internal/domain/user"
	"github.com/matthieutran/leafre-login/pkg/packet"
)

type SendCreateNewCharacter struct {
	Result    user.LoginResponse
	Character character.Character
}

var OpCodeCreateNewCharacterResult uint16 = 0xE

func WriteCreateNewCharacter(w io.Writer, send SendCreateNewCharacter) {
	pw := packet.NewPacketWriter()
	pw.WriteUInt16(OpCodeCreateNewCharacterResult)
	pw.WriteOne(byte(send.Result))
	if send.Result == user.LoginResponseSuccess {
		WriteCharacterStats(pw, send.Character)
		WriteCharacterLook(pw, send.Character)

		pw.WriteOne(0) // Rankings
		pw.WriteOne(0) // World ranking position
	} else {
		pw.WriteUInt32(0)
	}

	w.Write(pw.Packet())
}
