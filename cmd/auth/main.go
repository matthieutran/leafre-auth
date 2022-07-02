package main

import (
	"log"
	"os"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/matthieutran/leafre-auth/messaging"
	"github.com/matthieutran/leafre-auth/postgres"
)

var db *pgxpool.Pool

func main() {
	var wg sync.WaitGroup

	log.Println("Leafre - Authentication")

	wg.Add(1)
	s := messaging.Init(db, os.Getenv("NATS_URI"))
	defer s.Stop()

	wg.Wait()
}

func init() {
	var err error

	// Initialize db
	db, err = postgres.Init()
	if err != nil {
		log.Fatal("Error connecting to DB:", err)
	}
}
