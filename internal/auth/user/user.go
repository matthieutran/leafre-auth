package user

import (
	"encoding/json"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        int    `json:"id"`
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
}

type Users []User

type UserRepository interface {
	// Login validates the login details in the `UserForm` object and returns the user's id and error (where applicable).
	Login(UserForm) (id int, err error)
	// Register attempts to creates and store a new `User`. Returns the `User` object back with error (where applicable).
	Register(User) (User, error)
}

var ErrDBFail = errors.New("DB fail")

var ErrNotRegistered = errors.New("user not registered")
var ErrWrongPassword = errors.New("incorrect password")

var ErrEmailAlreadyRegistered = errors.New("user already registered")
var ErrUserAlreadyRegistered = errors.New("email already registered")

func hashPassword(password string) string {
	p, _ := bcrypt.GenerateFromPassword([]byte(password), 8)

	return string(p)
}

func comparePassword(password, hashed string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			// Password incorrect
			err = ErrWrongPassword
			return
		}

		// DB password is corrupt? - ErrHashTooShort
		log.Printf("Error comparing password from database... Stored hash corrupt (%s)?: %s", password, err)
		err = ErrDBFail
	}

	return
}
