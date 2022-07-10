package handler

import (
	"context"
	"io"
	"log"

	"github.com/matthieutran/leafre-login/internal/app/handler/reader"
	"github.com/matthieutran/leafre-login/internal/app/handler/writer"
	"github.com/matthieutran/leafre-login/internal/domain/character"
	"github.com/matthieutran/leafre-login/pkg/packet"
)

const OpCodeCheckDuplicatedID uint16 = 0x15

type HandlerCheckDuplicatedID struct {
	characterService character.CharacterService
}

func NewHandlerCheckDuplicatedID(characterService character.CharacterService) HandlerCheckDuplicatedID {
	return HandlerCheckDuplicatedID{characterService: characterService}
}

func (h *HandlerCheckDuplicatedID) Handle(w io.Writer, p packet.Packet) {
	recv := reader.ReadCheckDuplicatedID(p)
	duplicate, err := h.characterService.CheckName(context.Background(), recv.Name)
	if err != nil {
		log.Println("Error checking duplicate name:", err)
		return
	}

	send := writer.SendDuplicatedIDResult{
		Name:      recv.Name,
		Duplicate: duplicate,
	}
	writer.WriteCheckDuplicatedIDResult(w, send)
}

func (h *HandlerCheckDuplicatedID) String() string {
	return "CheckDuplicatedID"
}
