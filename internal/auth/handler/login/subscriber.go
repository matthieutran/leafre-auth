package login

import (
	"github.com/matthieutran/duey"
	"github.com/matthieutran/leafre-auth/internal/auth/user"
)

const subject = "auth.login"

func LoginSubscriber(s *duey.EventStreamer, userRepository user.UserRepository) func() (string, duey.Handler) {
	return func() (string, duey.Handler) {
		cb := func(_, reply string, f user.UserForm) {
			Login(s, reply, userRepository, f)
		}

		return subject, cb
	}
}
