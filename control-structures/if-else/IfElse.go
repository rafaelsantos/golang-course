package main

import (
	"fmt"
	"math/rand"
	"time"
)

func printResult(score float64) {
	if score >= 7 {
		fmt.Println("Approved", score)
	} else {
		fmt.Println("Reproved", score)
	}
}

func conceptScore(n float64) string {
	if n >= 9 && n <= 10 {
		return "A"
	} else if n >= 8 && n < 9 {
		return "B"
	} else if n >= 5 && n < 8 {
		return "C"
	} else if n >= 3 && n < 5 {
		return "D"
	} else {
		return "E"
	}
}

func randomNumer() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return r.Intn(10)
}

func main() {
	printResult(7)
	printResult(5)

	fmt.Println(conceptScore(10))
	fmt.Println(conceptScore(8))
	fmt.Println(conceptScore(7))
	fmt.Println(conceptScore(4))

	//if init
	if x := randomNumer(); x > 5 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOK")
	}
}
