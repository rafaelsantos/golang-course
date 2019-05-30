package main

import (
	"fmt"
	"log"
)

var sum = func(a, b int) int {
	return a + b
}

func main() {
	log.Println(sum(2, 3))

	subtract := func(a, b int) int {
		return a - b
	}

	fmt.Println(subtract(2, 3))
}
