package register

import (
	"github.com/matthieutran/duey"
	"github.com/matthieutran/leafre-auth/internal/auth/database"
)

type payload struct {
	Username string
	Password string
	Email    string
	Birthday string
}

const subjectSub = "auth.register"

func RegisterSubscriber(s *duey.EventStreamer, db *database.DB) func() (string, duey.Handler) {
	return func() (string, duey.Handler) {
		cb := func(_, reply string, p payload) {
			Register(s, reply, db, p)
		}

		return subjectSub, cb
	}
}
