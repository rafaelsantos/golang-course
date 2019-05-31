package main

import (
	"fmt"
	"log"
)

func factorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("Invalid number: %d", n)
	case n == 0:
		return 1, nil
	default:
		previous, _ := factorial(n - 1)
		return n * previous, nil
	}
}

func factorialUnsigned(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		return n * factorialUnsigned(n-1)
	}
}

func main() {
	result, _ := factorial(5)
	log.Println(result)

	_, error := factorial(-4)

	if error != nil {
		log.Println(error)
	}

	log.Println(factorialUnsigned(3))
}
