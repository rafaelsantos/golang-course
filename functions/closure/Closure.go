package main

import "log"

func closure() func() {
	x := 10
	var function = func() {
		log.Println(x * 10)
	}

	return function
}

func main() {
	x := 20
	log.Println(x)

	//It will print the value of x * 10 inside the closure scope
	print := closure()
	print()
}
