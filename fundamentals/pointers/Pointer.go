package main

import "fmt"

func main() {
	//Go has not arithmetic of pointers
	var p *int

	i := 1

	//Get memory address of i
	p = &i

	*p++ //get value of pointer

	fmt.Println(p, *p, i, &i)
}
