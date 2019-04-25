package main

import "fmt"

func buy(work1, work2 bool) (bool, bool, bool) {
	buyTV50 := work1 && work2
	buyTV32 := work1 != work2 //Exclusive or
	buyIceCream := work1 || work2

	return buyTV50, buyTV32, buyIceCream
}

func main() {
	TV50, TV32, icecream := buy(false, false)
	fmt.Printf("TV50: %t, TV32: %t, Ice cream: %t, Healthy: %t\n", TV50, TV32, icecream, !icecream)

	TV50, TV32, icecream = buy(true, true)
	fmt.Printf("TV50: %t, TV32: %t, Ice cream: %t, Healthy: %t", TV50, TV32, icecream, !icecream)
}
