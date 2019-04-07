package main

import "fmt"

func main() {
	var ide = "Visual Studio Code"

	//Type inferred by compiler
	var version = 1.330

	fmt.Println(ide, version)

	name, iteration, enable := "VSCode", 001, true
	fmt.Println(name, iteration, enable)

	var (
		path       = "/go"
		permission = 770
	)

	fmt.Println(path, permission)
}
