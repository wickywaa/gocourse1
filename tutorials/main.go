package main

import (
	"log"

	"github.com/wickywaa/gocourse1/helpers"
)

func main() {

	log.Println("hello")

	var myVar helpers.SomeType
	myVar.TypeName = "some name"
	log.Println(myVar.TypeName)

}
