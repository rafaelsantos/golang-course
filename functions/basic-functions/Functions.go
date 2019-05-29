package main

import (
	"fmt"
	"log"
)

func f1() {
	log.Println("A")
}

func f2(p string, f int) {
	log.Println(p, f)
}

func f3() string {
	return ""
}

func f4(i, x float64) string {
	return fmt.Sprint(i, x)
}

func f() (string, error) {
	return ".", nil
}

func main() {
	f1()
	f2("AB", 34)
	v, g := f3(), f4(4, 12.34)
	log.Println(v, g)

	g, k := f()
	log.Println(g, k)
}
