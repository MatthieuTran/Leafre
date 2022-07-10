package session

import "context"

// SessionService is a domain-level service acting as a facade for a SessionRepository
type SessionService interface {
	CreateSession(ctx context.Context, session Session) error
	GetSessionByID(ctx context.Context, id string) (Session, error)
	RemoveSession(ctx context.Context, id string) error
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
