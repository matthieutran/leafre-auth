package main

import (
	"log"
	"os"
	"sync"

	"github.com/matthieutran/leafre-auth/internal/auth/database"
	"github.com/matthieutran/leafre-auth/internal/auth/messaging"
)

func main() {
	var wg sync.WaitGroup

	log.Println("Leafre - Authentication")

	wg.Add(1)
	db, err := database.Init()
	if err != nil {
		log.Fatal("Error connecting to DB:", err)
	}

	wg.Add(1)
	s := messaging.Init(db, os.Getenv("NATS_URI"))
	defer s.Stop()

	wg.Wait()
}
