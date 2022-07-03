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

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResult struct {
	Code operation.LoginRequestCode `json:"code"`
	Id   int                        `json:"id"`
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

	var username, password string
	log.Print("Enter a Username: ")
	fmt.Scanln(&username)

	log.Printf("Enter a Password: ")
	fmt.Scanln(&password)

	res := &LoginResult{}
	err = s.Request("auth.login", &LoginPayload{
		Username: username,
		Password: password,
	}, res, 5*time.Second)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
	if res.Code == operation.Success {
		log.Println("Success!")
	}

}
