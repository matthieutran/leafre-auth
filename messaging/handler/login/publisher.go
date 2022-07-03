package login

import (
	"github.com/matthieutran/duey"
	"github.com/matthieutran/leafre-auth/pkg/operation"
)

func PublishLoginResponse(s *duey.EventStreamer, subject string, code operation.LoginRequestCode, id int) {
	s.Publish(subject, &response{Code: code, Id: id})
}
