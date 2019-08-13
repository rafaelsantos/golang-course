package main

import (
	"fmt"
	"time"
)

func routine(channel chan int) {
	fmt.Println("Running!")
	channel <- 1
	fmt.Println("Data written in channel: 1")
	channel <- 2
	fmt.Println("Data written in channel: 2")
	channel <- 3
	fmt.Println("Data written in channel: 3")
	channel <- 4
	fmt.Println("Data written in channel: 4")
	channel <- 5
	fmt.Println("Data written in channel: 5")
	channel <- 6
	fmt.Println("Data written in channel: 6")
}

func main() {
	//The channel will be blocked when the buffers limit
	//is reached and will be released only when the buffer
	//data is read
	channel := make(chan int, 3)
	go routine(channel)

	time.Sleep(time.Second)

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
