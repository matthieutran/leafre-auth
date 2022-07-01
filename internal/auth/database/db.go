package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

type DB struct {
	conn *pgxpool.Pool
}

func Init() (db *DB, err error) {
	db = &DB{}

	log.Println("Connecting to DB")
	dbURI := os.Getenv("POSTGRES_URI")
	conn, err := pgxpool.Connect(context.Background(), dbURI)
	if err != nil {
		return
	}

	db.conn = conn
	sql := `
	CREATE TABLE IF NOT EXISTS users(
		id            SERIAL       PRIMARY KEY,
		username      VARCHAR(32)  NOT NULL UNIQUE,
		password      VARCHAR(255) NOT NULL,
		email		  VARCHAR(32)  NOT NULL UNIQUE,
		created_at    TIMESTAMP    DEFAULT current_timestamp,
		updated_at    TIMESTAMP    DEFAULT current_timestamp
	);
	`
	if _, err = db.conn.Exec(context.Background(), sql); err != nil {
		return
	}

	log.Println("Successfully connected to DB")
	return
}

func (d *DB) Conn() *pgxpool.Pool {
	return d.conn
}

func (d *DB) Stop() {
	d.conn.Close()
}
