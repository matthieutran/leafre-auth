// The register tool is an interface for you to create an account on the server
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/matthieutran/duey"
	"github.com/matthieutran/leafre-auth/pkg/operation"
	"github.com/nats-io/nats.go"
)

type registerPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Result struct {
	Code int
}

type RegResult struct {
	Code operation.RegisterStatusCode
	Id   int
}

func main() {
	natsURI := os.Getenv("NATS_URI")
	if natsURI == "" {
		natsURI = nats.DefaultURL
	}

	log.Println("Leafre - Create an Account")

	s, err := duey.Init(natsURI)
	if err != nil {
		log.Fatal("Could not connect to messaging system:", err)
	}

	var username, password, email string
	fmt.Print("Enter a Username: ")
	fmt.Scanln(&username)

	fmt.Printf("Enter a Password: ")
	fmt.Scanln(&password)

	fmt.Printf("Enter an Email: ")
	fmt.Scanln(&email)

	res := &RegResult{}
	err = s.Request("auth.register", &registerPayload{
		Username: username,
		Password: password,
		Email:    email,
	}, res, 5*time.Second)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Result:", res)
	if res.Code == operation.RegisterSuccess {
		log.Println("Success!")
	}

}
