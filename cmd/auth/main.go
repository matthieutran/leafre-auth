package main

import (
	"log"
	"sync"

	"github.com/matthieutran/duey"
	"github.com/matthieutran/leafre-auth/internal/auth/handler/login"
)

type LoginPayload struct {
	Username string
	Password string
}

func main() {
	var wg sync.WaitGroup

	log.Println("Auth server")

	s, err := duey.Init()
	if err != nil {
		log.Fatal("Could not connect to messaging system:", err)
	}

	subscribers := []func() (string, duey.Handler){
		login.LoginSubscriber(s),
	}

	for _, subscriber := range subscribers {
		wg.Add(1)
		s.Subscribe(subscriber())
	}

	wg.Wait()
}
