package login

import (
	"github.com/matthieutran/duey"
	"github.com/matthieutran/leafre-auth/internal/auth/operation"
)

type loginResponse struct {
	Code operation.LoginRequestCode
	Id   int
}

func PublishLoginResponse(s *duey.EventStreamer, subject string, res loginResponse) {
	s.Publish(subject, res)
}
