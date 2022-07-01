package register

import (
	"errors"

	"github.com/matthieutran/duey"
	"github.com/matthieutran/leafre-auth/internal/auth/operation"
	"github.com/matthieutran/leafre-auth/internal/auth/user"
)

// Register attempts to create a new user
func Register(s *duey.EventStreamer, subject string, userRepository user.UserRepository, u user.User) {
	var code operation.RegisterStatusCode
	code = operation.RegisterSuccess

	newUser, err := userRepository.Register(u)
	if err != nil {
		if errors.Is(err, user.ErrUserAlreadyRegistered) {
			// Username is taken
			code = operation.RegisterDupeUsername
		} else if errors.Is(err, user.ErrEmailAlreadyRegistered) {
			// Email is taken
			code = operation.RegisterDupeEmail
		} else {
			// Some other weird DB error
			code = operation.RegisterServerError
		}
		return
	}

	PublishRegisterResponse(s, subject, code, newUser)
}
