package main

import (
	"fmt"
	"log"
)

type printable interface {
	toString() string
}

type sporty interface {
	turbo()
}

type person struct {
	name    string
	surname string
}

type product struct {
	name  string
	price float64
}

type bmw struct {
	model   string
	turboOn bool
	speed   int
}

//interfaces are implemented implicity
func (person person) toString() string {
	return person.name + " " + person.surname
}

func (product product) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", product.name, product.price)
}

func (bmw *bmw) turbo() {
	bmw.turboOn = true
}

func print(prt printable) {
	log.Println(prt.toString())
}

func main() {
	var any printable = person{"Johnny", "Wislow"}
	log.Println(any.toString())
	print(any)

	//Using the same var as type interface to instantiate another struct
	//Polymorphism
	any = product{"Phone", 2334.56}
	log.Println(any.toString())
	print(any)

	car1 := bmw{"B40", false, 0}
	car1.turbo()

	//The implementation of turbo() in bmw struct uses a pointer
	var car2 sporty = &bmw{"B40", false, 0}
	car2.turbo()

	log.Println(car1, car2)
}
