package handler

import (
	"context"
	"log"

	"github.com/matthieutran/leafre-login/internal/app/handler/reader"
	"github.com/matthieutran/leafre-login/internal/app/handler/writer"
	"github.com/matthieutran/leafre-login/internal/domain/session"
	"github.com/matthieutran/leafre-login/internal/domain/user"
	"github.com/matthieutran/leafre-login/pkg/packet"
)

const OpCodeCheckPassword uint16 = 0x1

type HandlerCheckPassword struct {
	authService    user.AuthService
	userService    user.UserService
	sessionService session.SessionService
}

func NewHandlerCheckPassword(authService user.AuthService, userService user.UserService, sessionService session.SessionService) HandlerCheckPassword {
	return HandlerCheckPassword{
		authService:    authService,
		userService:    userService,
		sessionService: sessionService,
	}
}

func (h *HandlerCheckPassword) Handle(s session.Session, p packet.Packet) {
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

	if res == user.LoginResponseSuccess {
		u, err := h.userService.GetUserByName(context.Background(), recv.Username)
		log.Println("USER ID:", u.ID)
		if err != nil {
			log.Printf("Could not get user object (%s): %s", recv.Username, err)
		}
		h.sessionService.SetSessionAccount(context.Background(), s.ID(), u)
	}

	// Write CheckPasswordResult
	writer.WriteCheckPasswordResult(s, send)
}

func (h *HandlerCheckPassword) String() string {
	return "CheckPassword"
}
