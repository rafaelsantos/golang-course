package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	score := 6.9
	_score := int(score)

	fmt.Println(_score)

	//Will print the value in ASCII code which represents 97
	fmt.Println("Number " + string(97))

	//Convert the number 97 to string
	fmt.Println(strconv.Itoa(97))

	//convert the string 97 to integer
	fmt.Println(strconv.Atoi("97"))

	result, _ := strconv.ParseBool("true")
	if result {
		fmt.Println(strconv.ParseBool("true"))
	}

}
