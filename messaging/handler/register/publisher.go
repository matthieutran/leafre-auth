package register

import (
	"github.com/matthieutran/duey"
	"github.com/matthieutran/leafre-auth/pkg/operation"
)

type response struct {
	Code operation.RegisterStatusCode `json:"code"`
	Id   int                          `json:"auth_id"`
}

func PublishRegisterResponse(s *duey.EventStreamer, subject string, code operation.RegisterStatusCode, id int) {
	res := &response{
		Code: code,
		Id:   id,
	}

	s.Publish(subject, res)
}
