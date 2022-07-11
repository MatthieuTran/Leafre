package handler

import (
	"github.com/matthieutran/leafre-login/internal/domain/session"
	"github.com/matthieutran/leafre-login/pkg/packet"
)

const OpCodeEnableSPWRequest uint16 = 0x5

type HandlerEnableSPWRequest struct {
}

func NewHandlerEnableSPWRequest() HandlerSelectWorld {
	return HandlerSelectWorld{}
}

func (h *HandlerEnableSPWRequest) Handle(s session.Session, p packet.Packet) {

}

func (h *HandlerEnableSPWRequest) String() string {
	return "EnableSPWRequest"
}
