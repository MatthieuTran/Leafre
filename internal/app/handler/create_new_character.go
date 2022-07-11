package handler

import (
	"context"
	"io"
	"log"

	"github.com/matthieutran/leafre-login/internal/app/handler/reader"
	"github.com/matthieutran/leafre-login/internal/app/handler/writer"
	"github.com/matthieutran/leafre-login/internal/domain/character"
	"github.com/matthieutran/leafre-login/internal/domain/user"
	"github.com/matthieutran/leafre-login/pkg/packet"
)

const OpCodeCreateNewCharacter uint16 = 0x16

type CreateNewCharacter struct {
	charService character.CharacterService
}

func NewCreateNewCharacter(characterService character.CharacterService) CreateNewCharacter {
	return CreateNewCharacter{charService: characterService}
}

func (h *CreateNewCharacter) Handle(w io.Writer, p packet.Packet) {
	recv := reader.ReadCreateNewCharacter(p)
	charDetails := character.CharacterForm{
		Name:      recv.Name,
		Job:       recv.Race,
		SubJob:    recv.SubJob,
		Face:      recv.Face,
		Hair:      recv.Hair,
		HairColor: recv.HairColor,
		Skin:      byte(recv.Skin),
		Coat:      recv.Coat,
		Pants:     recv.Pants,
		Shoes:     recv.Shoes,
		Weapon:    recv.Weapon,
		Gender:    recv.Gender,
	}

	char, err := h.charService.CreateCharacter(context.Background(), charDetails)
	if err != nil {
		log.Printf("Error creating character (name: %s): %s", recv.Name, err)
		return
	}

	result := user.LoginResponseSuccess
	send := writer.SendCreateNewCharacter{
		Result:    result,
		Character: char,
	}
	writer.WriteCreateNewCharacter(w, send)
}

func (h *CreateNewCharacter) String() string {
	return "CreateNewCharacter"
}
