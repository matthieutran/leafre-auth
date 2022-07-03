package login

import (
	"github.com/matthieutran/duey"
	auth "github.com/matthieutran/leafre-auth"
	"github.com/matthieutran/leafre-auth/pkg/operation"
)

const subject = "auth.login"

type response struct {
	Code operation.LoginRequestCode `json:"code"`
	Id   int                        `json:"id"`
}

func LoginSubscriber(s *duey.EventStreamer, userRepository auth.UserRepository) func() (string, duey.Handler) {
	return func() (string, duey.Handler) {
		cb := func(_, reply string, f auth.UserForm) {
			Login(s, reply, userRepository, f)
		}

		return subject, cb
	}
}
