package main

import "fmt"

func average(numbers ...float64) float64 {
	total := 0.0
	for _, number := range numbers {
		total += number
	}

	return total / float64(len(numbers))
}

func print(approved ...string) {
	for _, approved := range approved {
		fmt.Println(approved)
	}
}

func main() {
	fmt.Println(average(2.3, 5.6, 8.9, 3.4))

	approved := []string{"Johnny", "Gary"}
	print(approved...)
}
