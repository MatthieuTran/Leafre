package handler

import (
	"io"

	"github.com/matthieutran/duey"
	"github.com/matthieutran/packet"
)

const OpCodeWorldRequest uint16 = 0xB

type HandlerWorldRequest struct {
}

func (h *HandlerWorldRequest) Handle(w io.Writer, es *duey.EventStreamer, p packet.Packet) {

}

func (h *HandlerWorldRequest) String() string {
	return "WorldRequest"
}
