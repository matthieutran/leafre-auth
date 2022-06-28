package login

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/matthieutran/duey"
	"github.com/matthieutran/leafre-auth/internal/auth/database"
	"github.com/matthieutran/leafre-auth/internal/auth/operation"
	"golang.org/x/crypto/bcrypt"
)

// Login validates the given username and password combination
func Login(s *duey.EventStreamer, subject string, db *database.DB, username, password string) {
	var hashed string // Password to compare with

	// Fetch user
	err := db.Conn().QueryRow(context.Background(), `SELECT "password" FROM "users" WHERE "username"='$1'`, &username).Scan(&hashed)
	if err != nil {
		// No account registered
		if errors.Is(err, pgx.ErrNoRows) {
			PublishLoginResponse(s, subject, operation.NotRegistered)
			return
		}
		// Unknown error
		log.Println("Error validating login:", err)
		PublishLoginResponse(s, subject, operation.DBFail)
		return
	}

	// User successfully fetched, compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if err != nil {
		// Password incorrect
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			PublishLoginResponse(s, subject, operation.IncorrectPassword)
			return
		}
		// DB password is corrupt? - ErrHashTooShort
		log.Printf("Error comparing password from database for user %s... Password corrupt?: %s", username, err)
		PublishLoginResponse(s, subject, operation.DBFail)
		return
	}

	// Publish the login result
	PublishLoginResponse(s, subject, operation.Success)
}
