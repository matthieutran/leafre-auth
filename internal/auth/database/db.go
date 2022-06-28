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

func Init() (db DB, err error) {
	log.Println("Connecting to DB")
	dbURI := os.Getenv("POSTGRES_URI")
	if db.conn, err = pgxpool.Connect(context.Background(), dbURI); err != nil {
		return
	}

	sql := `
	CREATE TABLE IF NOT EXISTS users(
		id            SERIAL   PRIMARY KEY,
		username    VARCHAR(32) NOT NULL UNIQUE,
		password    VARCHAR(32) NOT NULL,
		created_at    TIMESTAMP DEFAULT current_timestamp,
		updated_at    TIMESTAMP DEFAULT current_timestamp
	);
	`
	if _, err = db.conn.Exec(context.Background(), sql); err != nil {
		return
	}

	log.Println("Successfully connected to DB")
	return
}

func (d *DB) Stop() {
	d.conn.Close()
}
