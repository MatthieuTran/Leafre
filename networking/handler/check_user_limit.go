package handler

import (
	"io"

	"github.com/matthieutran/duey"
	"github.com/matthieutran/leafre-login/networking/reader"
	"github.com/matthieutran/leafre-login/networking/writer"
	"github.com/matthieutran/packet"
)

const OpCodeCheckUserLimit uint16 = 0x6

type HandlerCheckUserLimit struct {
}

func NewHandlerCheckUserLimit() HandlerCheckUserLimit {
	return HandlerCheckUserLimit{}
}

func (h *HandlerCheckUserLimit) Handle(w io.Writer, es *duey.EventStreamer, p packet.Packet) {
	_ = reader.ReadCheckUserLimit(p) // TODO: handle this packet based on the client's specified world Id
	writer.WriteCheckUserLimitResult(w)
}

func (h *HandlerCheckUserLimit) String() string {
	return "CheckUserLimit"
}
