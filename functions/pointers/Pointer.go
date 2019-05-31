package main

import "log"

func inc1(n int) {
	n++
}

func inc2(n *int) {
	//* is used to defer (get the value)
	*n++
}

func main() {
	n := 1
	inc1(n) //by value
	log.Println(n)

	inc2(&n) //by reference
	log.Println(n)
}
