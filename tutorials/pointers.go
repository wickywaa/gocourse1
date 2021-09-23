package main

import "log"

func main() {

	var myString string
	myString = "green"

	log.Println("mystring is set to", myString)

	changeWithPointer(&myString)

	log.Println(myString)
}

func changeWithPointer(s *string) {

	newValue := "RED"
	*s = newValue

}
