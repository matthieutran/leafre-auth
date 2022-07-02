package postgres

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func Init() (conn *pgxpool.Pool, err error) {
	log.Println("Connecting to DB")
	dbURI := os.Getenv("POSTGRES_URI")
	conn, err = pgxpool.Connect(context.Background(), dbURI)
	if err != nil {
		return
	}

	sql := `CREATE TABLE IF NOT EXISTS users(
		id            SERIAL       PRIMARY KEY,
		username      VARCHAR(32)  NOT NULL UNIQUE,
		password      VARCHAR(255) NOT NULL,
		email		  VARCHAR(32)  NOT NULL UNIQUE,
		created_at    TIMESTAMP    DEFAULT current_timestamp,
		updated_at    TIMESTAMP    DEFAULT current_timestamp
	);
	`
	if _, err = conn.Exec(context.Background(), sql); err != nil {
		return
	}

	log.Println("Successfully connected to DB")
	return
}
