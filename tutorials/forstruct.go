package main

import "log"

func main() {

	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User

	users = append(users, User{"john", "smith", "j1234@gmail.com", 28})
	users = append(users, User{"geoff", "smongh", "johns3434rfmith@gmail.com", 100})
	users = append(users, User{"amy", "rtrfe", "jwewrfg@gmail.com", 32})
	users = append(users, User{"betty", "gatt", "eerwerf@gmail.com", 38})
	users = append(users, User{"mo", "gobo", "wertwertwe@gmail.com", 2})

	for _, person := range users {

		log.Println(person.FirstName, person.LastName, person.Email, person.Age)
	}
}
