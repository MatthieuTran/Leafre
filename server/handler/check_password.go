package handler

import (
	"io"

	"github.com/matthieutran/duey"
	login "github.com/matthieutran/leafre-login"
	"github.com/matthieutran/leafre-login/server/reader"
	"github.com/matthieutran/leafre-login/server/writer"
	"github.com/matthieutran/packet"
)

const OpCodeCheckPassword uint16 = 0x1

type HandlerCheckPassword struct {
	userRepository login.UserRepository
}

func NewHandlerCheckPassword(userRepository login.UserRepository) HandlerCheckPassword {
	return HandlerCheckPassword{
		userRepository: userRepository,
	}
}

func (h *HandlerCheckPassword) Handle(w io.Writer, es *duey.EventStreamer, p packet.Packet) {
	recv := reader.ReadLogin(p)
	user, code := h.userRepository.Login(
		login.UserForm{
			Username: recv.Username,
			Password: recv.Password,
		},
	)

	writer.WriteCheckPasswordResult(w, code, user)
}

func (h *HandlerCheckPassword) String() string {
	return "CheckPassword"
}
