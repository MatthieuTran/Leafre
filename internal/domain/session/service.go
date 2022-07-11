package session

import (
	"context"

	"github.com/matthieutran/leafre-login/internal/domain/user"
)

// SessionService is a domain-level service acting as a facade for a SessionRepository
type SessionService interface {
	CreateSession(ctx context.Context, session Session) error
	GetSessionByID(ctx context.Context, id string) (Session, error)
	RemoveSession(ctx context.Context, id string) error
	SetSessionAccount(ctx context.Context, id string, user user.User) (err error)
}

func NewSessionService(sr SessionRepository) SessionService {
	return &defaultSessionService{repository: sr}
}

type defaultSessionService struct {
	repository SessionRepository
}

func (ss *defaultSessionService) CreateSession(ctx context.Context, session Session) error {
	return ss.repository.Add(ctx, session)
}

func (ss *defaultSessionService) GetSessionByID(ctx context.Context, id string) (Session, error) {
	return ss.repository.GetByID(ctx, id)
}

func (ss *defaultSessionService) RemoveSession(ctx context.Context, id string) error {
	return ss.repository.Destroy(ctx, id)
}

func (ss *defaultSessionService) SetSessionAccount(ctx context.Context, id string, account user.User) (err error) {
	s, _ := ss.GetSessionByID(ctx, id)
	s.Account = account

	ss.repository.Update(ctx, s)
	return nil
}
