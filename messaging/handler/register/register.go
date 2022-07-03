package register

import (
	"errors"

	"github.com/matthieutran/duey"
	auth "github.com/matthieutran/leafre-auth"
	"github.com/matthieutran/leafre-auth/pkg/operation"
)

// Register attempts to create a new user
func Register(s *duey.EventStreamer, subject string, userRepository auth.UserRepository, form auth.UserForm) {
	var code operation.RegisterStatusCode
	code = operation.RegisterSuccess

	user := auth.User{
		Username: form.Username,
		Password: form.Password,
		Email:    form.Email,
	}

	id, err := userRepository.Register(user)
	if err != nil {
		if errors.Is(err, auth.ErrUserAlreadyRegistered) {
			// Username is taken
			code = operation.RegisterDupeUsername
		} else if errors.Is(err, auth.ErrEmailAlreadyRegistered) {
			// Email is taken
			code = operation.RegisterDupeEmail
		} else {
			// Some other weird DB error
			code = operation.RegisterServerError
		}
	}

	PublishRegisterResponse(s, subject, code, id)
}
