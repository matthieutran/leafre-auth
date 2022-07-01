package messaging

import (
	"log"

	"github.com/matthieutran/duey"
	"github.com/matthieutran/leafre-auth/internal/auth/database"
	"github.com/matthieutran/leafre-auth/internal/auth/handler/login"
	"github.com/matthieutran/leafre-auth/internal/auth/handler/register"
	"github.com/matthieutran/leafre-auth/internal/auth/user"
)

func Init(db *database.DB, uri string) *duey.EventStreamer {
	s, err := duey.Init(uri)
	if err != nil {
		log.Fatal("Could not connect to messaging system:", err)
	}

	userRepository := user.NewUserPostgresRepository()

	subscribers := []func() (string, duey.Handler){
		login.LoginSubscriber(s, userRepository),
		register.RegisterSubscriber(s, userRepository),
	}

	for _, subscriber := range subscribers {
		s.Subscribe(subscriber())
	}

	return s
}
