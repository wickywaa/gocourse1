package main

import "fmt"

func main() {

	fmt.Println("hello world")
	var i int
	i = 7

	fmt.Println(7 + i)

	var whatToSay string = "hello again"

	fmt.Println(whatToSay)

	whatwassaid, theotherting := saySomething("how are you", "I am great")

	fmt.Println(whatwassaid)
	fmt.Println(theotherting)

}

func saySomething(s string, t string) (string, string) {

	return s, t

}
