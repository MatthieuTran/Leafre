package handler

import (
	"io"

	"github.com/matthieutran/leafre-login/internal/app/handler/reader"
	"github.com/matthieutran/leafre-login/internal/app/handler/writer"
	"github.com/matthieutran/leafre-login/pkg/packet"
)

const OpCodeCheckUserLimit uint16 = 0x6

type HandlerCheckUserLimit struct {
}

func NewHandlerCheckUserLimit() HandlerCheckUserLimit {
	return HandlerCheckUserLimit{}
}

func (h *HandlerCheckUserLimit) Handle(w io.Writer, p packet.Packet) {
	_ = reader.ReadCheckUserLimit(p)

	send := writer.SendCheckUserLimit{}
	writer.WriteCheckUserLimitResult(w, send)
}

func (h *HandlerCheckUserLimit) String() string {
	return "CheckUserLimit"
}
