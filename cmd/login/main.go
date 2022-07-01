package main

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/matthieutran/duey"
	"github.com/matthieutran/leafre-auth/internal/auth/operation"
)

type LoginPayload struct {
	Username string
	Password string
}

func main() {
	var wg sync.WaitGroup

	log.Println("Auth server")

	s, err := duey.Init(os.Getenv("NATS_URI"))
	if err != nil {
		log.Fatal("Could not connect to messaging system:", err)
	}

	type payload struct {
		Username string
		Password string
		Email    string
		Birthday string
	}

	type Result struct {
		Code int
	}

	type RegResult struct {
		Code operation.RegisterStatusCode
		Id   int
	}

	res := &Result{}
	regResult := &RegResult{}
	err = s.Request("auth.login", &payload{Username: "matt", Password: "matt12"}, res, 5*time.Second)
	// err = s.Request("auth.register", &payload{Username: "matt2", Password: "matt12", Email: "matthieuktran3@gmail.com", Birthday: "1999-04-28"}, regResult, 5*time.Second)
	// "auth.register", &payload{Username: "matt", Password: "matt12", Email: "matthieuktran@gmail.com", Birthday: "1999-04-28"}, res, 5*time.Second)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res)
	log.Println(regResult)

	wg.Wait()
}
