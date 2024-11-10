package utils

import (
	"log"

	emailverifier "github.com/AfterShip/email-verifier"
)

func ValidateEmail(email string) bool {
	verifier := emailverifier.NewVerifier()

	result, err := verifier.Verify(email)
	if err != nil {
		log.Println("verify email address failed, error is: ", err)
		return false
	}

	if !result.Syntax.Valid {
		log.Println("email address syntax is invalid")
		return false
	}

	return result.Syntax.Valid
}