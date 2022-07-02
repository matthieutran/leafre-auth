package register

import (
	"github.com/matthieutran/duey"
	auth "github.com/matthieutran/leafre-auth"
	"github.com/matthieutran/leafre-auth/pkg/operation"
)

type response struct {
	Code operation.RegisterStatusCode `json:"code"`
	auth.User
}

func PublishRegisterResponse(s *duey.EventStreamer, subject string, code operation.RegisterStatusCode, u auth.User) {
	res := &response{
		Code: code,
		User: u,
	}

	s.Publish(subject, res)
}
