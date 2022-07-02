package register

import (
	"errors"

	"github.com/matthieutran/duey"
	auth "github.com/matthieutran/leafre-auth"
	"github.com/matthieutran/leafre-auth/pkg/operation"
)

// Register attempts to create a new user
func Register(s *duey.EventStreamer, subject string, userRepository auth.UserRepository, u auth.User) {
	var code operation.RegisterStatusCode
	code = operation.RegisterSuccess

	newUser, err := userRepository.Register(u)
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
		return
	}

	PublishRegisterResponse(s, subject, code, newUser)
}
