package main

import "fmt"

func change(p1, p2 int) (second int, first int) {
	second = p2
	first = p1

	//Clean return
	return
}

func main() {
	k, g := change(2, 3)
	fmt.Println(k, g)

	f, e := change(6, 8)
	fmt.Println(e, f)
}
