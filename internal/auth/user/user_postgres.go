package user

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v4"
	"github.com/matthieutran/leafre-auth/internal/auth/database"
)

type UserPostgresRepository struct {
	db *database.DB
}

func NewUserPostgresRepository() *UserPostgresRepository {
	return &UserPostgresRepository{}
}

func (r UserPostgresRepository) Login(form UserForm) (id int, err error) {
	var hashed string // Password to compare with
	id = -1           // Initialize id as -1

	// Fetch user
	query := `SELECT "id", "password" FROM "users" WHERE "username"=$1`
	err = r.db.Conn().QueryRow(context.Background(), query, &form.Username).Scan(&id, &hashed)
	if err != nil {
		// No account registered
		if errors.Is(err, pgx.ErrNoRows) {
			err = ErrNotRegistered
			return
		}
		// Unknown error
		log.Println("Error validating login:", err)
		err = ErrDBFail
		return
	}

	// User successfully fetched, compare passwords
	return id, comparePassword(form.Password, hashed)
}

func (r UserPostgresRepository) Register(u User) (user User, err error) {
	hashed := hashPassword(u.Password)

	// Check if username and email already exist
	if !r.checkUsername(u.Username) {
		err = ErrUserAlreadyRegistered
		return
	}
	if !r.checkEmail(u.Email) {
		err = ErrEmailAlreadyRegistered
		return
	}

	// Register user
	query := `INSERT INTO "users" ("username", "password", "email") VALUES ($1, $2, $3) RETURNING "id";`
	err = r.db.Conn().QueryRow(context.Background(), query, &u.Username, &hashed, &u.Email).Scan(&user.Id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			// We check earlier if already exists, but another user could have been made right before inserting
			if pgErr.Code == pgerrcode.UniqueViolation {
				err = ErrUserAlreadyRegistered
				return
			}
		}

		log.Println("Error creating account:", err)
		err = ErrDBFail
		return
	}

	log.Printf("Created account (User ID: %d)", user.Id)
	return
}

func (r UserPostgresRepository) checkUsername(username string) bool {
	var exists bool

	query := `SELECT EXISTS(SELECT 1 FROM "users" WHERE "username"=$1) as "exists";`
	r.db.Conn().QueryRow(context.Background(), query, &username).Scan(&exists) // Check if username already exists

	return !exists
}

func (r UserPostgresRepository) checkEmail(email string) bool {
	var exists bool

	query := `SELECT EXISTS(SELECT 1 FROM "users" WHERE "email"=$1) as "exists";`
	r.db.Conn().QueryRow(context.Background(), query, &email).Scan(&exists) // Check if email already exists

	return !exists
}
