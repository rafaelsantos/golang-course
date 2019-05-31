package main

import "fmt"

func fetchValue(number int) int {
	//Will print before the return of function
	defer fmt.Println("End")

	if number > 5000 {
		fmt.Println("Higher")
		return number
	}

	fmt.Println("Lower")
	return number
}

func main() {
	fmt.Println(fetchValue(7000))
}
