package user

import (
	"context"
	"errors"
)

// UserRepository provides an interface for accessing the data-layer
type UserRepository interface {
	Add(ctx context.Context, user User) error
	GetById(ctx context.Context, id int) (User, error)
	GetByName(ctx context.Context, name string) (User, error)
	Update(ctx context.Context, user User) error
	Destroy(ctx context.Context, id int) (err error)
}

var ErrAlreadyExists = errors.New("username already exists")
var ErrEmailAlreadyExists = errors.New("email already exists")
var ErrUserDoesNotExist = errors.New("user does not exist")
