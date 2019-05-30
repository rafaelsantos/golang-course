package main

import "log"

func calculation(a, b int) int {
	return a + b
}

func run(function func(int, int) int, p1, p2 int) int {
	return function(p1, p2)
}

func main() {
	result := run(calculation, 3, 4)
	log.Println(result)
}
