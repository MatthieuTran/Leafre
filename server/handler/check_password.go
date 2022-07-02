package handler

import (
	"io"

	"github.com/matthieutran/duey"
	"github.com/matthieutran/leafre-login/messaging/command"
	"github.com/matthieutran/leafre-login/server/reader"
	"github.com/matthieutran/leafre-login/server/writer"
	"github.com/matthieutran/packet"
)

const OpCodeCheckPassword uint16 = 0x1

type HandlerCheckPassword struct {
}

func (h *HandlerCheckPassword) Handle(w io.Writer, es *duey.EventStreamer, p packet.Packet) {
	payload := reader.ReadLogin(p)
	req := &command.RequestLogin{Username: payload.Username, Password: payload.Password} // Read packet and create LoginRequest struct
	res := command.CheckLogin(es, req)                                                   // Request login validation through event

	writer.WriteCheckPasswordResult(w, res.Code, res.Id, payload.Username)
}

func (h *HandlerCheckPassword) String() string {
	return "CheckPassword"
}
