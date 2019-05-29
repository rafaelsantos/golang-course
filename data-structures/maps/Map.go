package main

import "log"

func main() {
	//maps should be initialized
	approved := make(map[int]string)

	approved[12345678] = "Paul"
	approved[34567345] = "Gary"
	approved[81823467] = "Joana"

	log.Println(approved)

	for code, name := range approved {
		log.Printf("%s (Code: %d)\n", name, code)
	}

	log.Println(approved[34567345])
	delete(approved, 34567345)

	log.Println(approved)

	customers := map[string]float64{
		"Aaron": 2345.67,
	}

	customers["Sue"] = 3456.78

	log.Println(customers)

	customersByName := map[string]map[string]float64{
		"Robert": {
			"Robert Jonas": 4567.89,
			"Robert Crow":  2345.67,
		},
		"Gary": {
			"Gary Wisdom": 2312.56,
		},
	}

	delete(customersByName, "Robert")
	log.Println(customersByName)
}
