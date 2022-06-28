package register

import (
	"github.com/matthieutran/duey"
	"github.com/matthieutran/leafre-auth/internal/auth/operation"
)

type response struct {
	Code operation.RegisterStatusCode
	Id   int
}

func PublishRegisterResponse(s *duey.EventStreamer, subject string, code operation.RegisterStatusCode, id int) {
	res := &response{
		Code: code,
		Id:   id,
	}

	s.Publish(subject, res)
}
