package register

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/matthieutran/duey"
	"github.com/matthieutran/leafre-auth/internal/auth/database"
	"github.com/matthieutran/leafre-auth/internal/auth/operation"
	"golang.org/x/crypto/bcrypt"
)

// Register attempts to create a new user
func Register(s *duey.EventStreamer, subject string, db *database.DB, p payload) {
	var err error
	hashed := hashPassword(p.Password)

	// Check if username and email already exist
	if !checkUsername(db, p) {
		PublishRegisterResponse(s, subject, operation.RegisterDupeUsername, -1)
		return
	}
	if !checkEmail(db, p) {
		PublishRegisterResponse(s, subject, operation.RegisterDupeEmail, -1)
		return
	}

	// Register user
	var id int // ID of new user
	err = db.Conn().QueryRow(context.Background(), `INSERT INTO "users" ("username", "password", "email", "birthday") VALUES ($1, $2, $3, $4) RETURNING "id";`, &p.Username, &hashed, &p.Email, &p.Birthday).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			// We check earlier if already exists, but another user could have been made right before inserting
			if pgErr.Code == pgerrcode.UniqueViolation {
				PublishRegisterResponse(s, subject, operation.RegisterDupeUsername, -1)
				return
			}
		}
		log.Println("Error creating account:", err)
		PublishRegisterResponse(s, subject, operation.RegisterServerError, -1)
		return
	}

	log.Printf("Created account (User ID: %d)", id)
	PublishRegisterResponse(s, subject, operation.RegisterSuccess, id)
}

func checkUsername(db *database.DB, p payload) bool {
	var exists bool
	db.Conn().QueryRow(context.Background(), `SELECT EXISTS(SELECT 1 FROM "users" WHERE "username"=$1) as "exists";`, &p.Username).Scan(&exists) // Check if username already exists

	return !exists
}

func checkEmail(db *database.DB, p payload) bool {
	var exists bool
	db.Conn().QueryRow(context.Background(), `SELECT EXISTS(SELECT 1 FROM "users" WHERE "email"=$1) as "exists";`, &p.Email).Scan(&exists) // Check if email already exists

	return !exists
}

func hashPassword(password string) string {
	p, _ := bcrypt.GenerateFromPassword([]byte(password), 8)

	return string(p)
}
