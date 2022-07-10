package handler

import (
	"io"

	"github.com/matthieutran/leafre-login/internal/app/handler/reader"
	"github.com/matthieutran/leafre-login/internal/app/handler/writer"
	"github.com/matthieutran/leafre-login/internal/domain/character"
	"github.com/matthieutran/leafre-login/internal/domain/user"
	"github.com/matthieutran/leafre-login/pkg/packet"
)

const OpCodeCreateNewCharacter uint16 = 0x16

type CreateNewCharacter struct {
	characterService character.CharacterService
}

func NewCreateNewCharacter(characterService character.CharacterService) CreateNewCharacter {
	return CreateNewCharacter{characterService: characterService}
}

func (h *CreateNewCharacter) Handle(w io.Writer, p packet.Packet) {
	_ = reader.ReadCreateNewCharacter(p)

	result := user.LoginResponseSuccess
	send := writer.SendCreateNewCharacter{
		Result: result,
	}
	writer.WriteCreateNewCharacter(w, send)
}

func (h *CreateNewCharacter) String() string {
	return "CreateNewCharacter"
}
