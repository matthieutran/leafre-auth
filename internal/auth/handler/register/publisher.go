package register

import (
	"github.com/matthieutran/duey"
	"github.com/matthieutran/leafre-auth/internal/auth/operation"
	"github.com/matthieutran/leafre-auth/internal/auth/user"
)

type response struct {
	Code operation.RegisterStatusCode `json:"code"`
	user.User
}

func PublishRegisterResponse(s *duey.EventStreamer, subject string, code operation.RegisterStatusCode, u user.User) {
	res := &response{
		Code: code,
		User: u,
	}

	s.Publish(subject, res)
}
