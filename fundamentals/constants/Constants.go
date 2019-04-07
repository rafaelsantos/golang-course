package main

import (
	"fmt"
	calc "math"
)

func main() {
	const PI float64 = 3.1415

	const (
		FLAG   = "RESTRICTED"
		NUMBER = 4.56
	)

	fmt.Println(calc.Pow(2.34, PI), FLAG, NUMBER)
}
