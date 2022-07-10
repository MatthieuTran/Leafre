package service

import (
	"github.com/matthieutran/leafre-login/internal/adapters/inmem"
	"github.com/matthieutran/leafre-login/internal/domain/session"
	"github.com/matthieutran/leafre-login/internal/domain/user"
)

type Application struct {
	SessionService              session.SessionService
	SessionCommunicationService session.SessionCommunicationService
	AuthService                 user.AuthService
}

func NewApplication() *Application {
	sr := inmem.NewSessionRepository()                // Create Session Repository
	ss := session.NewSessionService(sr)               // Create Session Service and inject repository
	scs := session.NewSessionCommunicationService(sr) // Create Session Communication Service and inject repository
	ur := inmem.NewUserRepository()                   // Create User Repository
	us := user.NewAuthService(ur)                     // Create User Auth Service and inject user repository

	return &Application{
		SessionService:              ss,
		SessionCommunicationService: scs,
		AuthService:                 us,
	}
}
