package world

import (
	"context"
	"errors"
)

// WorldRepository provides an interface for accessing the data-layer
type WorldRepository interface {
	GetAll(ctx context.Context) (Worlds, error)
	GetByID(ctx context.Context, id byte) (World, error)
}

var ErrDoesNotExist = errors.New("world does not exist")
