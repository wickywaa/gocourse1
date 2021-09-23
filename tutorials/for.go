package main

import "log"

func main() {

	for i := 0; i <= 10000; i++ {

		log.Println(i)
	}
	animals := []string{"dog", "fish", "horse", "cow", "cat"}

	rangeslices(animals)

	createmap(animals)

}

func rangeslices(slice []string) {

	for i, name := range slice {

		log.Println(i, name)
	}

}

func createmap(slice []string) {

	myMap := make(map[int]string)

	for i, name := range slice {

		myMap[i] = name
		log.Println(myMap)
		log.Println(slice)
	}

}
