package main

import (
	"encoding/json"
	"log"
	"strings"
)

type data struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

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

type person struct {
	name    string
	surname string
}

//Working with pseudo inheritance
type car struct {
	name  string
	speed int
}

type ferrari struct {
	car   //anonymous field
	turbo bool
}

//Customizing types
type score float64

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

func (person person) getName() string {
	return person.name + " " + person.surname
}

func (person *person) setName(name string) {
	values := strings.Split(name, " ")
	person.name = values[0]
	person.surname = values[1]
}

func (score score) between(start, end float64) bool {
	return float64(score) >= start && float64(score) <= end
}

func scoreResult(score score) string {
	if score.between(9.0, 10.0) {
		return "A"
	} else if score.between(7.0, 8.99) {
		return "B"
	}

	return "C"
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

	person1 := person{"Carlos", "Rodney"}
	log.Println(person1.getName())

	person1.setName("Juliana Lemes")
	log.Println(person1.getName())

	//Working with pseudo inheritance
	//Special type of composition
	car := ferrari{}

	//The type car declared in ferrari struct will generate anonymous fields for ferrari
	car.name = "F40"
	car.speed = 0
	car.turbo = true

	log.Println(car.name, car.turbo)
	log.Println(car)

	//Using custom types
	log.Println(scoreResult(9.4))

	//Convert to JSON
	d1 := data{1, "Notebook", 1899.90, []string{"Eletronic", "Computer"}}
	js, error := json.Marshal(d1)

	if error != nil {
		log.Println(error)
	}

	log.Println(string(js))

	var d2 data
	value := `{"id":2,"name":"Pencil","price":7.9,"tags":["Util","Draw"]}`
	error = json.Unmarshal([]byte(value), &d2)

	log.Println(d2.Name)
}
