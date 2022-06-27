package login

import (
	"log"

	"github.com/matthieutran/duey"
)

type payload struct {
	Username string
	Password string
}

const subject = "auth.login"

func LoginSubscriber(s *duey.EventStreamer) func() (string, duey.Handler) {
	return func() (string, duey.Handler) {
		cb := func(_, reply string, p payload) {
			log.Println(reply, p)
			doLogin(s, reply, p.Username, p.Password)
		}

		return subject, cb
	}
}
