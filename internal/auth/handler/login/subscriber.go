package login

import (
	"github.com/matthieutran/duey"
	"github.com/matthieutran/leafre-auth/internal/auth/database"
)

type payload struct {
	Username string
	Password string
}

const subject = "auth.login"

func LoginSubscriber(s *duey.EventStreamer, db *database.DB) func() (string, duey.Handler) {
	return func() (string, duey.Handler) {
		cb := func(_, reply string, p payload) {
			Login(s, reply, db, p.Username, p.Password)
		}

		return subject, cb
	}
}
