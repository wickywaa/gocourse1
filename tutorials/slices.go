package main

import "log"

func main() {

	names := []string{"Gav", "Dom", "Finley", "Connor"}

	log.Println(names)

	var myString []string

	myString = append(myString, "hello")
	myString = append(myString, "how")
	myString = append(myString, "are")
	myString = append(myString, "you")

	log.Println(myString)
	log.Println(names[1:2])
}
