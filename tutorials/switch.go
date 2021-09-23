package main

import "log"

func main() {

	myVar := "dog"

	switch myVar {

	case "cat":
		log.Println("cat is cat")

	case "dog":
		log.Println("case is dog")

	default:
		log.Println("not cat or dog , must be some other freak ")

	}
}
