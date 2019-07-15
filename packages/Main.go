package main

import "fmt"

func main() {
	p1 := Point{2.0, 5.6}
	p2 := Point{5.6, 7.8}

	fmt.Println(cathetus(p1, p2))
	fmt.Println(Distance(p1, p2))
}
