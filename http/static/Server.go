package main

import (
	"log"
	"net/http"
)

func main() {
	fserver := http.FileServer(http.Dir("public"))
	http.Handle("/", fserver)

	log.Println("Running...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
