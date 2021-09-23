package main

import "log"

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {

	var myVar myStruct

	myVar.FirstName = "John"

	myVar2 := myStruct{
		FirstName: "mary",
	}
	log.Println(myVar.printFirstName(), myVar2.printFirstName())

}
