package main

import (
	"log"
	"sync"

	"github.com/matthieutran/leafre-auth/internal/auth/messaging"
)

func main() {
	var wg sync.WaitGroup

	log.Println("Leafre - Authentication")

	wg.Add(1)
	messaging.Init()

	wg.Wait()
}
