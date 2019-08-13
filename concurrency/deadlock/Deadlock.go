package main

import (
	"fmt"
	"time"
)

func routine(channel chan int) {
	time.Sleep(2 * time.Second)

	//Blocking operation
	//Write data in channel
	channel <- 1
	fmt.Println("Writing data in channel")
}

func main() {
	channel := make(chan int) //No buffer

	go routine(channel)

	//Reading data from channel
	//Wait until some data is written to the channel
	fmt.Println(<-channel)
	fmt.Println("Data readed")

	//Deadlock : no data available for channel
	//All go routines are dead
	fmt.Println(<-channel)
}
