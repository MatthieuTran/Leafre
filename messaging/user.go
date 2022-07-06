package messaging

import (
	"errors"

	"github.com/matthieutran/duey"
	"github.com/matthieutran/leafre-login/messaging/command"
	"github.com/matthieutran/leafre-login/user"
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

type UserService struct {
	es *duey.EventStreamer
}

func NewUserService(es *duey.EventStreamer) (r user.Service) {
	r.es = es

	return
}

// Login validates the login details in the `UserForm` object and returns the user's object and error (where applicable)
func (r UserService) Login(form user.LoginForm) (u user.User, code user.LoginResponse) {
	var authUser AuthUser
	res := command.CheckLogin(r.es, form) // Request login validation through event
	if res.Code == user.LoginResponseSuccess {
		// TODO: change to call GetUser() and return instead.
		authUser = AuthUser{
			AuthId:   res.Id,
			Username: form.Username,
			Password: form.Password,
		}

		profileUser, _ := r.getProfileById(res.Id)
		u = r.stitchUsers(authUser, profileUser)
	}

	return u, res.Code
}

// stitchUsers combines an auth user and a profile user to fulfill the generic User model
func (r UserService) stitchUsers(authUser AuthUser, profileUser ProfileUser) user.User {
	return user.User{
		Id:       authUser.AuthId,
		Username: authUser.Username,
		Password: authUser.Password,
		Email:    authUser.Email,
		Gender:   profileUser.Gender,
	}
}

func (r UserService) getAuthById(int) (user AuthUser, err error) {
	return AuthUser{}, errors.New("not implemented yet")
}

func (r UserService) getProfileById(int) (ProfileUser, error) {
	return ProfileUser{Gender: 0}, nil
}

//GetById fetches a user by its ID
func (r UserService) GetById(id int) (user.User, error) {
	authUser, err := r.getAuthById(id)
	if err != nil {
		return user.User{}, err
	}
	profileUser, _ := r.getProfileById(id)

	return r.stitchUsers(authUser, profileUser), err
}
