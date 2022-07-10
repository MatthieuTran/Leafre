package handler

import (
	"io"

	"github.com/matthieutran/leafre-login/internal/app/handler/reader"
	"github.com/matthieutran/leafre-login/internal/app/handler/writer"
	"github.com/matthieutran/leafre-login/internal/domain"
	"github.com/matthieutran/leafre-login/internal/domain/user"
	"github.com/matthieutran/leafre-login/pkg/packet"
)

const OpCodeSelectWorld uint16 = 0x5

type HandlerSelectWorld struct {
	worldChannelService domain.WorldChannelService
}

func NewHandlerSelectWorld(worldChannelService domain.WorldChannelService) HandlerSelectWorld {
	return HandlerSelectWorld{worldChannelService: worldChannelService}
}

func (h *HandlerSelectWorld) Handle(w io.Writer, p packet.Packet) {
	_ = reader.ReadSelectWorld(p)

	result := user.LoginResponseSuccess
	send := writer.SendSelectWorld{
		Result: result,
	}
	writer.WriteSelectWorldResult(w, send)
}

func (h *HandlerSelectWorld) String() string {
	return "SelectWorld"
}
