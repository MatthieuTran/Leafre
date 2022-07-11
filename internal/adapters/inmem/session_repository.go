package inmem

import (
	"context"

	"github.com/matthieutran/leafre-login/internal/domain/session"
)

func NewSessionRepository() session.SessionRepository {
	sessions := make(map[string]session.Session)

	return &SessionRepository{sessions: sessions}
}

// SessionRepository implements SessionRepository with an in-memory map
type SessionRepository struct {
	sessions map[string]session.Session
}

// Add adds a session to the repository
func (sr *SessionRepository) Add(ctx context.Context, s session.Session) error {
	if _, exists := sr.sessions[s.ID()]; exists {
		return session.ErrAlreadyExists
	}

	sr.sessions[s.ID()] = s
	return nil
}

func (sr *SessionRepository) GetByID(ctx context.Context, id string) (s session.Session, err error) {
	s, exists := sr.sessions[id]
	if !exists {
		return s, session.ErrDoesNotExist
	}

	return
}

func (sr *SessionRepository) Update(ctx context.Context, s session.Session) error {
	sr.sessions[s.ID()] = s

	return nil
}

func (sr *SessionRepository) Destroy(ctx context.Context, id string) error {
	if _, exists := sr.sessions[id]; !exists {
		return session.ErrDoesNotExist
	}

	delete(sr.sessions, id)
	return nil
}
