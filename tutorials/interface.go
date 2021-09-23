package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {

	dog := Dog{
		Name:  "Samson",
		Breed: "German Shepherd",
	}

	gorilla := Gorilla{
		Name:          "Jock",
		Color:         "Grey",
		NumberOfTeeth: 38,
	}
	PrintInfo(&gorilla)
	PrintInfo(&dog)

}

func PrintInfo(a Animal) {
	fmt.Println("this animal says ", a.Says(), "and has ", a.NumberOfLegs(), "legs")
}

func (d *Dog) Says() string {
	return "woof"
}

func (d *Dog) NumberOfLegs() int {

	return 4
}

func (g *Gorilla) Says() string {

	return "oooh oohh ah ah "
}

func (g *Gorilla) NumberOfLegs() int {

	return 2
}
