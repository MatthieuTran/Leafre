package handler

import (
	"context"
	"log"

	"github.com/matthieutran/leafre-login/internal/app/handler/reader"
	"github.com/matthieutran/leafre-login/internal/app/handler/writer"
	"github.com/matthieutran/leafre-login/internal/domain"
	"github.com/matthieutran/leafre-login/internal/domain/character"
	"github.com/matthieutran/leafre-login/internal/domain/session"
	"github.com/matthieutran/leafre-login/internal/domain/user"
	"github.com/matthieutran/leafre-login/pkg/packet"
)

const OpCodeSelectWorld uint16 = 0x5

type HandlerSelectWorld struct {
	worldChannelService domain.WorldChannelService
	charService         character.CharacterService
}

func NewHandlerSelectWorld(worldChannelService domain.WorldChannelService, charService character.CharacterService) HandlerSelectWorld {
	return HandlerSelectWorld{worldChannelService: worldChannelService, charService: charService}
}

func (h *HandlerSelectWorld) Handle(s session.Session, p packet.Packet) {
	_ = reader.ReadSelectWorld(p)

	chars, err := h.charService.GetCharactersByAccount(context.Background(), s.Account.ID)
	if err != nil {
		log.Printf("Could not get characters for (accID: %d): %s", s.Account.ID, err)
	}

	result := user.LoginResponseSuccess
	send := writer.SendSelectWorld{
		Result:     result,
		Characters: chars,
	}

	writer.WriteSelectWorldResult(s, send)
}

func (h *HandlerSelectWorld) String() string {
	return "SelectWorld"
}
