package main

import (
	"log"
)

func main() {
	var isTrue bool
	isTrue = false

	if isTrue == true {
		log.Println("is true is ", isTrue)
	} else {

		log.Println("isTrue is  not true", isTrue)

	}

	log.Println(isCat("diog"))
	log.Println(checkNum(80, 150))

}

func isCat(cat string) bool {

	if cat == "cat" {

		return true
	} else {
		return false
	}
}

func checkNum(x int, y int) (bool, string) {

	if x > 99 && y < 200 {
		return true, "that is correct"
	} else {
		return false, "that is not correct"
	}
}
