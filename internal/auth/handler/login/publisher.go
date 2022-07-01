package login

import (
	"github.com/matthieutran/duey"
	"github.com/matthieutran/leafre-auth/internal/auth/operation"
)

type response struct {
	Code operation.LoginRequestCode `json:"code"`
	Id   int                        `json:"id"`
}

func PublishLoginResponse(s *duey.EventStreamer, subject string, code operation.LoginRequestCode, id int) {
	s.Publish(subject, &response{Code: code, Id: id})
}
