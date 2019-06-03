package main

import "log"

type product struct {
	name     string
	price    float64
	discount float64
}

//Creates a receiver of type product
func (product product) priceWithDiscount() float64 {
	return product.price * (1 - product.discount)
}

func main() {
	var product1 product
	product1 = product{
		name:     "Pencil",
		price:    1.79,
		discount: 0.04,
	}

	log.Println(product1.priceWithDiscount())
}
