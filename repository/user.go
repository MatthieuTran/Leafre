package repository

import (
	"errors"

	"github.com/matthieutran/duey"
	login "github.com/matthieutran/leafre-login"
	"github.com/matthieutran/leafre-login/messaging/command"
	"github.com/matthieutran/leafre-login/pkg/operation"
)

// An AuthUser is an instance of the user object that we get back from the auth service
type AuthUser struct {
	AuthId    int    `json:"auth_id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// A ProfileUser is an instance of the user object that we get back from the profile service
type ProfileUser struct {
	ProfileId int  `json:"profile_id"`
	Gender    byte `json:"gender"`
}

type UserRepository struct {
	es *duey.EventStreamer
}

func NewUserRepository(es *duey.EventStreamer) (r UserRepository) {
	r.es = es

	return
}

// Login validates the login details in the `UserForm` object and returns the user's object and error (where applicable)
func (r UserRepository) Login(form login.UserForm) (user login.User, code operation.CodeLoginRequest) {
	var authUser AuthUser
	res := command.CheckLogin(r.es, form) // Request login validation through event
	if res.Code == operation.LoginRequestSuccess {
		// TODO: change to call GetUser() and return instead.
		authUser = AuthUser{
			AuthId:   res.Id,
			Username: form.Username,
			Password: form.Password,
		}

		profileUser, _ := r.getProfileById(res.Id)
		user = r.stitchUsers(authUser, profileUser)
	}

	return user, res.Code
}

// stitchUsers combines an auth user and a profile user to fulfill the generic User model
func (r UserRepository) stitchUsers(authUser AuthUser, profileUser ProfileUser) login.User {
	return login.User{
		Id:       authUser.AuthId,
		Username: authUser.Username,
		Password: authUser.Password,
		Email:    authUser.Email,
		Gender:   profileUser.Gender,
	}
}

func (r UserRepository) getAuthById(int) (user AuthUser, err error) {
	return AuthUser{}, errors.New("not implemented yet")
}

func (r UserRepository) getProfileById(int) (ProfileUser, error) {
	return ProfileUser{Gender: 0}, nil
}

//GetById fetches a user by its ID
func (r UserRepository) GetById(id int) (login.User, error) {
	authUser, err := r.getAuthById(id)
	if err != nil {
		return login.User{}, err
	}
	profileUser, _ := r.getProfileById(id)

	return r.stitchUsers(authUser, profileUser), err
}
