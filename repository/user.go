package repository

import (
	auth "github.com/matthieutran/leafre-auth"
	"github.com/matthieutran/leafre-auth/crypto"
)

type UserModel interface {
	GetByUsername(string) (auth.User, error)
	Add(auth.User) (auth.User, error)
}

type UserRepository struct {
	user UserModel
}

func NewUserPostgresRepository(user UserModel) *UserRepository {
	return &UserRepository{user: user}
}

func (r UserRepository) Login(form auth.UserForm) (id int, err error) {
	user, err := r.user.GetByUsername(form.Username)
	if err != nil {
		return -1, err
	}

	// User successfully fetched, compare passwords
	return id, crypto.ComparePassword(form.Password, user.Password)
}

func (r UserRepository) Register(u auth.User) (user auth.User, err error) {
	// Hash password before storing to DB
	user.Password = crypto.HashPassword(u.Password)

	// Return added object (with ID) and error (where applicable)
	return r.user.Add(user)
}
