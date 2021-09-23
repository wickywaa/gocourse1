package main

import (
	"log"

	"github.com/wickywaa/gocourse1/helpers"
)

const numpool = 10000000

func calculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numpool)
	intChan <- randomNumber
}

func main() {

	intChan := make(chan int)
	defer close(intChan)

	go calculateValue(intChan)

	num := <-intChan
	log.Println(num)

}
