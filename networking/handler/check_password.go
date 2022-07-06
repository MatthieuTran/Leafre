package handler

import (
	"io"

	"github.com/matthieutran/duey"
	"github.com/matthieutran/leafre-login/networking/reader"
	"github.com/matthieutran/leafre-login/networking/writer"
	"github.com/matthieutran/leafre-login/user"
	"github.com/matthieutran/packet"
)

const OpCodeCheckPassword uint16 = 0x1

type HandlerCheckPassword struct {
	userService user.Service
}

func NewHandlerCheckPassword(userService user.Service) HandlerCheckPassword {
	return HandlerCheckPassword{
		userService: userService,
	}
}

func (h *HandlerCheckPassword) Handle(w io.Writer, es *duey.EventStreamer, p packet.Packet) {
	recv := reader.ReadLogin(p)
	user, code := h.userService.Login(
		user.LoginForm{
			Username: recv.Username,
			Password: recv.Password,
		},
	)

	writer.WriteCheckPasswordResult(w, code, user)
}

func (h *HandlerCheckPassword) String() string {
	return "CheckPassword"
}
