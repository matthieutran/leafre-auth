package register

import (
	"github.com/matthieutran/duey"
	"github.com/matthieutran/leafre-auth/internal/auth/user"
)

const subjectSub = "auth.register"

func RegisterSubscriber(s *duey.EventStreamer, userRepository user.UserRepository) func() (string, duey.Handler) {
	return func() (string, duey.Handler) {
		cb := func(_, reply string, p user.User) {
			Register(s, reply, userRepository, p)
		}

		return subjectSub, cb
	}
}
