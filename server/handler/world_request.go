package handler

import (
	"github.com/matthieutran/duey"
	"github.com/matthieutran/packet"
	"github.com/matthieutran/tcpserve"
)

const OpCodeWorldRequest uint16 = 0xB

type HandlerWorldRequest struct {
}

func (h *HandlerWorldRequest) Handle(*tcpserve.Session, *duey.EventStreamer, packet.Packet) {

}

func (h *HandlerWorldRequest) String() string {
	return "WorldRequest"
}
