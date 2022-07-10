package service

import (
	"github.com/matthieutran/leafre-login/internal/adapters/inmem"
	"github.com/matthieutran/leafre-login/internal/app/handler"
	"github.com/matthieutran/leafre-login/internal/domain"
	"github.com/matthieutran/leafre-login/internal/domain/session"
	"github.com/matthieutran/leafre-login/internal/domain/user"
)

type Application struct {
	SessionService session.SessionService
	AuthService    user.AuthService
	Handlers       map[uint16]handler.PacketHandler
}

func NewApplication() *Application {
	// Repositories
	channelRepo := inmem.NewChannelRepository()
	sessionRepo := inmem.NewSessionRepository()
	userRepo := inmem.NewUserRepository()
	worldRepo := inmem.NewWorldRepository()

	// Services
	authService := user.NewAuthService(userRepo)
	sessionService := session.NewSessionService(sessionRepo)
	worldChannelService := domain.NewWorldChannelService(worldRepo, channelRepo)

	// handlers is a map of opcodes to packet handlers
	handlers := make(map[uint16]handler.PacketHandler)
	addHandler := func(opcode uint16, h handler.PacketHandler) {
		handlers[opcode] = h
	}

	// Initialize packet handlers
	checkPassword := handler.NewHandlerCheckPassword(authService)
	worldRequest := handler.NewHandlerWorldRequest(worldChannelService)
	checkUserLimit := handler.NewHandlerCheckUserLimit()
	selectWorld := handler.NewHandlerSelectWorld(worldChannelService)

	// Add packet handlers to the map
	addHandler(handler.OpCodeCheckPassword, &checkPassword)
	addHandler(handler.OpCodeWorldRequest, &worldRequest)
	addHandler(handler.OpCodeCheckUserLimit, &checkUserLimit)
	addHandler(handler.OpCodeSelectWorld, &selectWorld)

	return &Application{
		SessionService: sessionService,
		AuthService:    authService,

		Handlers: handlers,
	}
}
