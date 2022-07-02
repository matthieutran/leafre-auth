package crypto

import (
	"errors"
	"log"

	auth "github.com/matthieutran/leafre-auth"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	p, _ := bcrypt.GenerateFromPassword([]byte(password), 8)

	return string(p)
}

func ComparePassword(password, hashed string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			// Password incorrect
			err = auth.ErrWrongPassword
			return
		}

		// DB password is corrupt? - ErrHashTooShort
		log.Printf("Error comparing password from database... Stored hash corrupt (%s)?: %s", password, err)
		err = auth.ErrDBFail
	}

	return
}
