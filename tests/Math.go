package tests

import (
	"fmt"
	"log"
	"strconv"
)

//Average : calculate the average of list of numbers
func Average(numbers ...float64) float64 {
	total := 0.0

	for _, number := range numbers {
		total += number
	}

	average := total / float64(len(numbers))
	roundAverage, error := strconv.ParseFloat(fmt.Sprintf("%.2f", average), 64)

	if error != nil {
		log.Println(error)
	}

	return roundAverage
}
