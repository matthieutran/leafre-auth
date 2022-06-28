package messaging

import (
	"log"

	"github.com/matthieutran/duey"
	"github.com/matthieutran/leafre-auth/internal/auth/database"
	"github.com/matthieutran/leafre-auth/internal/auth/handler/login"
	"github.com/matthieutran/leafre-auth/internal/auth/handler/register"
)

func Init(uri string, db *database.DB) *duey.EventStreamer {
	s, err := duey.Init(uri)
	if err != nil {
		log.Fatal("Could not connect to messaging system:", err)
	}

	subscribers := []func() (string, duey.Handler){
		login.LoginSubscriber(s, db),
		register.RegisterSubscriber(s, db),
	}

	for _, subscriber := range subscribers {
		s.Subscribe(subscriber())
	}

	return s
}
