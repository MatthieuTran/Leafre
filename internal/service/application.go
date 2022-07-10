package service

import (
	"github.com/matthieutran/leafre-login/internal/adapters/inmem"
	"github.com/matthieutran/leafre-login/internal/app/handler"
	"github.com/matthieutran/leafre-login/internal/domain/session"
	"github.com/matthieutran/leafre-login/internal/domain/user"
)

type Application struct {
	SessionService              session.SessionService
	SessionCommunicationService session.SessionCommunicationService
	AuthService                 user.AuthService
	Handlers                    map[uint16]handler.PacketHandler
}

func NewApplication() *Application {
	// Repositories
	sr := inmem.NewSessionRepository() // Create Session Repository
	ur := inmem.NewUserRepository()    // Create User Repository

	// Services
	sessionService := session.NewSessionService(sr)                           // Create Session Service and inject repository
	sessionCommunicationService := session.NewSessionCommunicationService(sr) // Create Session Communication Service and inject repository
	authService := user.NewAuthService(ur)                                    // Create User Auth Service and inject user repository

	// handlers is a map of opcodes to packet handlers
	handlers := make(map[uint16]handler.PacketHandler)
	addHandler := func(opcode uint16, h handler.PacketHandler) {
		handlers[opcode] = h
	}

	// Initialize packet handlers
	checkPassword := handler.NewHandlerCheckPassword(authService)

	// Add packet handlers to the map
	addHandler(handler.OpCodeCheckPassword, &checkPassword)

	return &Application{
		SessionService:              sessionService,
		SessionCommunicationService: sessionCommunicationService,
		AuthService:                 authService,

		Handlers: handlers,
	}
}
