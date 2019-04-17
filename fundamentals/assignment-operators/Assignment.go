package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 3 //inference type
	i += 3
	i -= 3
	i /= 1
	i *= 1
	i %= 1

	fmt.Println(i)

	a, j, c := 2, 5, 7
	fmt.Println(a, c, j)

	a, c = j, a
	fmt.Println(a, j, c)
}
