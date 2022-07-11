package service

import (
	"github.com/matthieutran/leafre-login/internal/adapters/inmem"
	"github.com/matthieutran/leafre-login/internal/app/handler"
	"github.com/matthieutran/leafre-login/internal/domain"
	"github.com/matthieutran/leafre-login/internal/domain/character"
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
	charRepo := inmem.NewCharacterRepository()
	itemRepo := inmem.NewItemRepository()
	sessionRepo := inmem.NewSessionRepository()
	userRepo := inmem.NewUserRepository()
	worldRepo := inmem.NewWorldRepository()

	// Services
	authService := user.NewAuthService(userRepo)
	characterService := character.NewCharacterService(charRepo, itemRepo)
	sessionService := session.NewSessionService(sessionRepo)
	userService := user.NewUserService(userRepo)
	worldChannelService := domain.NewWorldChannelService(worldRepo, channelRepo)

	// handlers is a map of opcodes to packet handlers
	handlers := make(map[uint16]handler.PacketHandler)
	addHandler := func(opcode uint16, h handler.PacketHandler) {
		handlers[opcode] = h
	}

	// Initialize packet handlers
	checkDuplicatedID := handler.NewHandlerCheckDuplicatedID(characterService)
	checkPassword := handler.NewHandlerCheckPassword(authService, userService, sessionService)
	checkUserLimit := handler.NewHandlerCheckUserLimit()
	createNewCharacter := handler.NewHandlerCreateNewCharacter(characterService)
	clientDumpLog := handler.NewHandlerClientDumpLog()
	selectWorld := handler.NewHandlerSelectWorld(worldChannelService, characterService)
	worldRequest := handler.NewHandlerWorldRequest(worldChannelService)

	// Add packet handlers to the map
	addHandler(handler.OpCodeCheckPassword, &checkPassword)
	addHandler(handler.OpCodeCheckDuplicatedID, &checkDuplicatedID)
	addHandler(handler.OpCodeCheckUserLimit, &checkUserLimit)
	addHandler(handler.OpCodeCreateNewCharacter, &createNewCharacter)
	addHandler(handler.OpCodeWorldRequest, &worldRequest)
	addHandler(handler.OpCodeSelectWorld, &selectWorld)
	addHandler(handler.OpCodeClientDumpLog, &clientDumpLog)

	return &Application{
		SessionService: sessionService,
		AuthService:    authService,

		Handlers: handlers,
	}
}
