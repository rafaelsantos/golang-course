package main

import "fmt"

func main() {
	var a int
	var b float64
	var c bool
	var d string
	var e *int

	//String %v can be used to print any value type in console
	fmt.Printf("%v %v %v %v %v", a, b, c, d, e)

	//String %q can print nullable strings in console
	fmt.Printf("%v %v %v %q %v", a, b, c, d, e)
}
