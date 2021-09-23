package main

import (
	"errors"
	"log"
)

func main() {

	result, err := divide(100.0, 0.0)
	if err != nil {

		log.Println("there iis an error", err)
		return
	}

	log.Println("result of divison is ", result)

}

func divide(x, y float32) (float32, error) {

	var result float32

	if y == 0 {
		return result, errors.New("cannot divide by 0")
	}

	result = x / y
	return result, nil
}
