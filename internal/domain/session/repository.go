package session

import (
	"context"
	"errors"
)

// SessionRepository manages and keeps track of all the sessions
type SessionRepository interface {
	Add(ctx context.Context, s Session) error
	GetByID(ctx context.Context, id string) (Session, error)
	Destroy(ctx context.Context, id string) (err error)
}

var ErrAlreadyExists = errors.New("session already exists")
var ErrDoesNotExist = errors.New("session does not exist")
