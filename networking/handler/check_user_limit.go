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
	_ = reader.ReadCheckUserLimit(p)

	var bWarningLevel, bPopulateLevel byte // TODO: implement bWarningLevel and bPopulateLevel
	writer.WriteCheckUserLimitResult(w, bWarningLevel, bPopulateLevel)
}

func (h *HandlerCheckUserLimit) String() string {
	return "CheckUserLimit"
}
