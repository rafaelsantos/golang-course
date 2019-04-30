package main

import (
	"fmt"
	"time"
)

func conceptScore(n float64) string {
	var score = int(n)
	switch score {
	//Default behavior of switch is break in any of case block
	//Using fallthrough it will continue entering in others case blocks
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	default:
		return "K"
	}
}

func genericSwitch(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	}

	return "invalid"
}

func main() {
	fmt.Println(conceptScore(10))
	fmt.Println(conceptScore(8))
	fmt.Println(conceptScore(4))

	t := time.Now()

	//switch true
	//Get the first true case or default
	switch {
	case t.Hour() < 12:
		fmt.Println("12")
	default:
		fmt.Println(t.Hour())
	}

	fmt.Println(genericSwitch(12))
	fmt.Println(genericSwitch("12"))
}
