package login

import (
	"github.com/matthieutran/duey"
	"github.com/matthieutran/leafre-auth/internal/auth/operation"
)

func doLogin(s *duey.EventStreamer, subject, username, password string) {
	PublishLoginResponse(s, subject, operation.AuthFail)
}
