package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func (m *User) joinNames() string {

	fullname := make(map[string]string)
	fullname["fullName"] = m.FirstName + " " + m.LastName
	return fullname["fullName"]
}

func main() {

	dave := User{

		FirstName: "Gavin",
		LastName:  "Newton",
	}

	myMap := make(map[string]int)

	myMap["first"] = 2

	log.Println(myMap["first"])

	log.Println(dave.joinNames())

}
