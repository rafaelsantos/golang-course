package main

import "log"

//Will be executed as the first action in main function
func init() {
	log.Println("Go!")
}

func main() {
	log.Println("Main")
}
