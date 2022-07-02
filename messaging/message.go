package messaging

import (
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/matthieutran/duey"
	"github.com/matthieutran/leafre-auth/internal/auth/handler/login"
	"github.com/matthieutran/leafre-auth/internal/auth/handler/register"
	"github.com/matthieutran/leafre-auth/postgres"
	"github.com/matthieutran/leafre-auth/repository"
)

func Init(conn *pgxpool.Pool, uri string) *duey.EventStreamer {
	s, err := duey.Init(uri)
	if err != nil {
		log.Fatal("Could not connect to messaging system:", err)
	}

	// Inject database into model
	userModel := postgres.NewUserModel(conn)

	// Inject model into repository
	userRepository := repository.NewUserPostgresRepository(userModel)

	// Map subscriber handlers
	subscribers := []func() (string, duey.Handler){
		// Each subscriber is injected with whatever it needs
		login.LoginSubscriber(s, userRepository),
		register.RegisterSubscriber(s, userRepository),
	}

	// Add subscribers
	for _, subscriber := range subscribers {
		s.Subscribe(subscriber())
	}

	return s
}
