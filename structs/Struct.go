package main

import "log"

type product struct {
	name     string
	price    float64
	discount float64
}

type item struct {
	productID int
	quantity  int
	price     float64
}

type order struct {
	userID int
	itens  []item
}

//Creates a receiver of type product
func (product product) priceWithDiscount() float64 {
	return product.price * (1 - product.discount)
}

func (order order) getTotal() float64 {
	total := 0.0

	for _, item := range order.itens {
		total += item.price * float64(item.quantity)
	}

	return total
}

func main() {
	var product1 product
	product1 = product{
		name:     "Pencil",
		price:    1.79,
		discount: 0.04,
	}

	log.Println(product1.priceWithDiscount())

	order := order{
		userID: 1,
		itens: []item{
			item{1, 2, 12.10},
			item{2, 1, 23.49},
			item{11, 100, 3.24},
		},
	}

	log.Println(order.getTotal())
}
