package writer

import (
	"bytes"
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
	var charStats bytes.Buffer
	var charLook bytes.Buffer
	WriteCharacterStats(&charStats, send.Character)
	WriteCharacterLook(&charLook, send.Character)

	pw := packet.NewPacketWriter()
	pw.WriteUInt16(OpCodeCreateNewCharacterResult)
	pw.WriteOne(byte(send.Result))
	if send.Result == user.LoginResponseSuccess {
		pw.WriteBytes(charStats.Bytes())
		pw.WriteBytes(charLook.Bytes())

		pw.WriteOne(0) // Rankings
		pw.WriteOne(0) // World ranking position
	} else {
		pw.WriteUInt32(0)
	}

	w.Write(pw.Packet())
}
