package main

import (
	"fmt"
	"time"
)

var s = "seven"

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {

	user := User{

		FirstName: "Gavin",
		LastName:  "Newton",
	}

	fmt.Println(user.FirstName)
	fmt.Println(user)

}
