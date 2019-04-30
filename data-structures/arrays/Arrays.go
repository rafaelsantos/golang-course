package main

import "fmt"

func main() {
	var scores [4]float64
	fmt.Println(scores)

	scores[0], scores[1], scores[2] = 5.6, 7.2, 4.6
	fmt.Println(scores)

	total := 0.0

	for i := 0; i < len(scores); i++ {
		total += scores[i]
	}

	avg := total / float64(len(scores))

	fmt.Println(avg)
}
