package user

import (
	"context"
)

type UserService interface {
	GetUserByID(ctx context.Context, userID int) (u User, err error)
	GetUserByName(ctx context.Context, username string) (u User, err error)
}

func NewUserService(ur UserRepository) UserService {
	return defaultUserService{userRepo: ur}
}

type defaultUserService struct {
	userRepo UserRepository
}

func (s defaultUserService) GetUserByID(ctx context.Context, userID int) (u User, err error) {
	u, err = s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return
	}

	return
}

func (s defaultUserService) GetUserByName(ctx context.Context, username string) (u User, err error) {
	u, err = s.userRepo.GetByName(ctx, username)
	if err != nil {
		return
	}

	return
}
