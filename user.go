package login

import (
	"github.com/matthieutran/leafre-login/pkg/operation"
)

type User struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Gender    byte   `json:"gender"`
}

type Users []User

type UserForm struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	MachineId string `json:"machine_id"`
}

type UserRepository interface {
	// Login validates the login details in the `UserForm` object and returns the user's object and error (where applicable)
	Login(UserForm) (user User, code operation.CodeLoginRequest)

	//GetById fetches a user by its ID
	GetById(int) (user User, err error)
}
