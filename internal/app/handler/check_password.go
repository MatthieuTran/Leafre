package handler

import (
	"context"
	"io"

	"github.com/matthieutran/leafre-login/internal/app/handler/reader"
	"github.com/matthieutran/leafre-login/internal/app/handler/writer"
	"github.com/matthieutran/leafre-login/internal/domain/user"
	"github.com/matthieutran/leafre-login/pkg/packet"
)

const OpCodeCheckPassword uint16 = 0x1

type HandlerCheckPassword struct {
	authService user.AuthService
}

func NewHandlerCheckPassword(userService user.AuthService) HandlerCheckPassword {
	return HandlerCheckPassword{
		authService: userService,
	}
}

func (h *HandlerCheckPassword) Handle(w io.Writer, p packet.Packet) {
	// Read CheckPassword
	recv := reader.ReadLogin(p)
	form := user.AuthForm{
		Username: recv.Username,
		Password: recv.Password,
	}

	// Authenticate
	res := h.authService.Login(context.Background(), form)
	send := writer.SendCheckLogin{
		Result:         byte(res),
		Name:           recv.Username,
		NumOfCharacter: 4,
		V44:            1,
	}

	// Write CheckPasswordResult
	writer.WriteCheckPasswordResult(w, send)
}

func (h *HandlerCheckPassword) String() string {
	return "CheckPassword"
}
