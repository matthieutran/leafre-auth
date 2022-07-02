package login

import (
	"github.com/matthieutran/duey"
	auth "github.com/matthieutran/leafre-auth"
)

const subject = "auth.login"

func LoginSubscriber(s *duey.EventStreamer, userRepository auth.UserRepository) func() (string, duey.Handler) {
	return func() (string, duey.Handler) {
		cb := func(_, reply string, f auth.UserForm) {
			Login(s, reply, userRepository, f)
		}

		return subject, cb
	}
}
