package auth

// An AuthUser is an instance of the user object that we get back from the auth service
type AuthUser struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// A ProfileUser is an instance of the user object that we get back from the profile service
type ProfileUser struct {
	Id     int  `json:"id"`
	Gender byte `json:"gender"`
}

type User struct {
	AuthUser
	ProfileUser
}

type Users []User

type UserForm struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	MachineId string `json:"machine_id"`
}

type Users []User

type UserRepository interface {
	// Login validates the login details in the `UserForm` object and returns the user's object and error (where applicable)
	Login(UserForm) (user User, err error)

	//GetById fetches a user by its ID
	GetById(int) (user User, err error)
}
