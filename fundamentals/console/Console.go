package main

import (
	"fmt"
	calc "math"
)

func main() {

	var x = 2345 * calc.Pow(12.45, 56.78)
	y := fmt.Sprint(x)

	fmt.Println("Value of y", y)
	fmt.Printf("Value of y %.2f\n", x)

	w := 12
	k := 45.67
	var j = false
	param := "console"

	fmt.Printf("%d %.2f %t %s\n", w, k, j, param)
	fmt.Printf("%v %v %v %v", k, j, param, w)
}
