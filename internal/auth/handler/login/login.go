package login

import (
	"errors"

	"github.com/matthieutran/duey"
	"github.com/matthieutran/leafre-auth/internal/auth/operation"
	"github.com/matthieutran/leafre-auth/internal/auth/user"
)

// Login validates the given username and password combination
func Login(s *duey.EventStreamer, subject string, userRepo user.UserRepository, form user.UserForm) {
	var code operation.LoginRequestCode
	code = operation.Success

	id, err := userRepo.Login(form)
	if err != nil {
		if errors.Is(err, user.ErrNotRegistered) {
			// User does not exist in the storage
			code = operation.NotRegistered
		} else {
			// Some other weird DB error
			code = operation.DBFail
		}
	}

	PublishLoginResponse(s, subject, code, id)
}
