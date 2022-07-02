package postgres

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	auth "github.com/matthieutran/leafre-auth"
)

type UserModel struct {
	conn *pgxpool.Pool
}

func NewUserModel(conn *pgxpool.Pool) *UserModel {
	return &UserModel{conn: conn}
}

func (u *UserModel) ExistsUsername(username string) (exists bool) {
	query := `SELECT EXISTS(SELECT 1 FROM "users" WHERE "username"=$1) as "exists";`
	u.conn.QueryRow(context.Background(), query, &username).Scan(&exists) // Check if username already exists

	return !exists
}

func (u *UserModel) ExistsEmail(email string) (exists bool) {
	query := `SELECT EXISTS(SELECT 1 FROM "users" WHERE "email"=$1) as "exists";`
	u.conn.QueryRow(context.Background(), query, &email).Scan(&exists) // Check if email already exists

	return !exists
}

// GetByUsername fetches a user from the database
//
// Possible errors:
// auth.ErrNotRegistered -> User has not been registered under the provided username
// auth.ErrDBFail -> Query could not run because of some DB fetching error
func (u *UserModel) GetByUsername(username string) (user auth.User, err error) {
	user.Id = -1

	// Fetch user
	query := `SELECT "id", "password" FROM "users" WHERE "username"=$1`
	err = u.conn.QueryRow(context.Background(), query, username).Scan(&user.Id, &user.Password)
	if err != nil {
		// No account registered
		if errors.Is(err, pgx.ErrNoRows) {
			err = auth.ErrNotRegistered
			return
		}
		// Unknown error
		log.Println("Error fetching username:", err)
		err = auth.ErrDBFail
		return
	}

	// User successfully fetched, compare passwords
	return
}

// Add inserts the provided user into the database
//
// Possible Errors:
// auth.ErrUserAlreadyRegistered - Username has been taken
// auth.ErrEmailAlreadyRegistered - Email has been taken
// auth.ErrDBFail - Query could not run because of some DB fetching error
func (u *UserModel) Add(user auth.User) (auth.User, error) {
	var err error

	// Check if username exists
	if !u.ExistsUsername(user.Username) {
		err = auth.ErrUserAlreadyRegistered
		return user, err
	}
	// Check if email exists
	if !u.ExistsEmail(user.Email) {
		err = auth.ErrEmailAlreadyRegistered
		return user, err
	}

	// Register user
	query := `INSERT INTO "users" ("username", "password", "email") VALUES ($1, $2, $3) RETURNING "id";`
	err = u.conn.QueryRow(context.Background(), query, &user.Username, &user.Password, &user.Email).Scan(&user.Id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			// We check earlier if already exists, but another user could have been made right before inserting
			if pgErr.Code == pgerrcode.UniqueViolation {
				err = auth.ErrUserAlreadyRegistered
				return user, err
			}
		}

		log.Println("Error creating user:", err)
		err = auth.ErrDBFail
	}

	return user, err
}
