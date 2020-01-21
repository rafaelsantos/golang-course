package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func hour(writer http.ResponseWriter, request *http.Request) {
	value := time.Now().Format("02/01/2006 03:04:05")
	fmt.Fprintf(writer, "<h1>Hour: %s</h1>", value)
}

func main() {
	http.HandleFunc("/hour", hour)

	log.Println("Running...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
