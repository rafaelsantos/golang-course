package main

import "fmt"

type sporty interface {
	turboOn()
}

type luxurious interface {
	balize()
}

//Composition
type sportyLuxurious interface {
	sporty
	luxurious
	//...more methods
}

type bmw7 struct{}

func (b bmw7) turboOn() {
	fmt.Println("Turbo")
}

func (b bmw7) balize() {
	fmt.Println("Balize")
}

func main() {
	var b sportyLuxurious = bmw7{}
	b.turboOn()
	b.balize()
}
