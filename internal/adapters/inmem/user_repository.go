package inmem

import (
	"context"
	"time"

	"github.com/matthieutran/leafre-login/internal/domain/user"
)

// UserRepository implements `user.UserRepository` with an in-memory map
type UserRepository struct {
	users map[int]user.User
}

func NewUserRepository() user.UserRepository {
	// Create map of user id -> user
	users := make(map[int]user.User)
	r := &UserRepository{users: users}

	// Store initial data
	dummy_user := user.User{
		Name:     "matt",
		Password: "matt12",
		Email:    "matthieuktran@gmail.com",
		Birthday: time.Date(1999, time.April, 28, 0, 0, 0, 0, nil),
		Gender:   0,
	}
	r.Add(context.Background(), dummy_user)

	return &UserRepository{users: users}
}

func (r UserRepository) nameExists(name string) bool {
	for _, u := range r.users {
		if u.Name == name {
			return true
		}
	}

	return false
}

func (r UserRepository) Add(ctx context.Context, u user.User) error {
	if r.nameExists(u.Name) {
		return user.ErrAlreadyExists
	}

	if r.nameExists(u.Email) {
		return user.ErrEmailAlreadyExists
	}

	u.ID = len(r.users)
	r.users[u.ID] = u
	return nil
}

func (r UserRepository) GetById(ctx context.Context, id int) (u user.User, err error) {
	u, exists := r.users[id]
	if !exists {
		err = user.ErrEmailAlreadyExists
	}

	return
}

func (r UserRepository) GetByName(ctx context.Context, name string) (u user.User, err error) {
	for _, u := range r.users {
		if u.Name == name {
			return u, nil
		}
	}

	return u, user.ErrUserDoesNotExist
}

func (r UserRepository) Update(ctx context.Context, user user.User) error {
	return nil
}

func (r UserRepository) Destroy(ctx context.Context, id int) (err error) {
	delete(r.users, id)
	return nil
}
