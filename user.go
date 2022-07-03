package auth

import (
	"encoding/json"
	"errors"
)

type User struct {
	Id        int    `json:"auth_id"`
	Username  string `json:"username"`
	Password  string `json:"password,omitempty"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (u User) MarshalJSON() ([]byte, error) {
	type user User
	x := user(u)
	x.Password = ""

	return json.Marshal(x)
}

type UserForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Users []User

type UserRepository interface {
	// Login validates the login details in the `UserForm` object and returns the user's id and error (where applicable).
	Login(UserForm) (id int, err error)
	// Register attempts to creates and store a new `User` and returns the user's id and error (where applicable).
	Register(User) (id int, err error)
}

var ErrDBFail = errors.New("DB fail")

var ErrNotRegistered = errors.New("user not registered")
var ErrWrongPassword = errors.New("incorrect password")

var ErrEmailAlreadyRegistered = errors.New("user already registered")
var ErrUserAlreadyRegistered = errors.New("email already registered")
