package user

import (
	"context"
	"errors"
)

type AuthForm struct {
	Username string
	Password string
}

type AuthService interface {
	Login(ctx context.Context, user AuthForm) error
	Register(ctx context.Context, user AuthForm) (User, error)
}

var ErrIncorrectPassword = errors.New("incorrect password")

func NewAuthService(ur UserRepository) AuthService {
	return defaultAuthService{userRepo: ur}
}

type defaultAuthService struct {
	userRepo UserRepository
}

func (s defaultAuthService) Login(ctx context.Context, user AuthForm) (err error) {
	u, err := s.userRepo.GetByName(ctx, user.Username)
	if err != nil {
		return
	}

	if u.Password != user.Password {
		return ErrIncorrectPassword
	}

	return nil
}
func (s defaultAuthService) Register(ctx context.Context, user AuthForm) (u User, err error) {
	return
}
