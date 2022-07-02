package register

import (
	"github.com/matthieutran/duey"
	auth "github.com/matthieutran/leafre-auth"
)

const subjectSub = "auth.register"

func RegisterSubscriber(s *duey.EventStreamer, userRepository auth.UserRepository) func() (string, duey.Handler) {
	return func() (string, duey.Handler) {
		cb := func(_, reply string, p auth.User) {
			Register(s, reply, userRepository, p)
		}

		return subjectSub, cb
	}
}
