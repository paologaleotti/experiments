/// newtype rust pattern in go, test
/// Note: email must be in separate package to avoid use of private fields

package main

import (
	"fmt"
	"log"
	"newtype/email"
)

func getUserByEmail(email email.Email) {
	// get user by email
}

func main() {
	rawEmail := "hello@lmao.com"
	validatedEmail, err := email.EmailFromString(rawEmail)
	if err != nil {
		log.Fatal(err)
	}

	getUserByEmail(rawEmail) // This will not compile

	getUserByEmail(validatedEmail) // This is ok

	fmt.Println(validatedEmail.String())
}
