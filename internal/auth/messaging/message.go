package messaging

import (
	"log"

	"github.com/matthieutran/duey"
	"github.com/matthieutran/leafre-auth/internal/auth/handler/login"
)

func Init(uri string) *duey.EventStreamer {
	s, err := duey.Init(uri)
	if err != nil {
		log.Fatal("Could not connect to messaging system:", err)
	}

	subscribers := []func() (string, duey.Handler){
		login.LoginSubscriber(s),
	}

	for _, subscriber := range subscribers {
		s.Subscribe(subscriber())
	}

	return s
}
