package login

import (
	"github.com/matthieutran/duey"
	"github.com/matthieutran/leafre-auth/internal/auth/operation"
)

type response struct {
	Code operation.LoginRequestCode
}

func PublishLoginResponse(s *duey.EventStreamer, subject string, code operation.LoginRequestCode) {
	res := &response{
		Code: code,
	}

	s.Publish(subject, res)
}
