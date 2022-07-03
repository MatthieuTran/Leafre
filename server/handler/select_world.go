package handler

import (
	"io"

	"github.com/matthieutran/duey"
	"github.com/matthieutran/leafre-login/pkg/operation"
	"github.com/matthieutran/leafre-login/server/reader"
	"github.com/matthieutran/leafre-login/server/writer"
	"github.com/matthieutran/packet"
)

const OpCodeSelectWorld uint16 = 0x5

type HandlerSelectWorld struct {
}

func NewHandlerSelectWorld() HandlerSelectWorld {
	return HandlerSelectWorld{}
}

func (h *HandlerSelectWorld) Handle(w io.Writer, es *duey.EventStreamer, p packet.Packet) {
	_ = reader.ReadSelectWorld(p)

	result := operation.LoginRequestSuccess
	writer.WriteSelectWorldResult(w, result)
}

func (h *HandlerSelectWorld) String() string {
	return "SelectWorld"
}
