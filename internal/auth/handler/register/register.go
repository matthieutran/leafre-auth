package register

import (
	"context"
	"log"

	"github.com/matthieutran/duey"
	"github.com/matthieutran/leafre-auth/internal/auth/database"
)

// Register attempts to create a new user
func Register(s *duey.EventStreamer, subject string, db *database.DB, p payload) {
	var id int

	// Register user and store its id. No need to check if user already exists. we will handle through errors
	err := db.Conn().QueryRow(context.Background(), `INSERT INTO "users" ("username", "password", "email", "birthday") VALUES ($1, $2, $3, $4) RETURNING "id";`, &p.Username, &p.Password, &p.Email, &p.Birthday).Scan(&id)
	if err != nil {
		log.Println("Error creating account:", err)
	}

	log.Printf("Created account (User ID: %d)", id)
}
