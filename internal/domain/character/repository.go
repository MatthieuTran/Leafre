package character

import (
	"context"
	"errors"
)

// CharacterRepository provides an interface for accessing the data-layer
type CharacterRepository interface {
	Add(ctx context.Context, character Character) (uint32, error)
	GetByID(ctx context.Context, id uint32) (Character, error)
	GetByAccountID(ctx context.Context, accountId uint32) (Characters, error)
	GetByName(ctx context.Context, name string) (Character, error)
	Update(ctx context.Context, character Character) error
	Destroy(ctx context.Context, id uint32) (err error)
}

var ErrAlreadyExists = errors.New("character name already exists")
var ErrCharDoesNotExist = errors.New("character name provided does not exist")
