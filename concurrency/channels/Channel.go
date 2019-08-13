package main

import (
	"fmt"
	"time"
)

func isPrime(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}

func primes(number int, channel chan int) {
	start := 2
	for i := 0; i < number; i++ {
		for prime := start; ; prime++ {
			if isPrime(prime) {
				channel <- prime
				start = prime + 1
				time.Sleep(time.Millisecond * 100)
				break
			}
		}
	}
	//Finish the communication in channel
	//No data will be sent to channel
	//Used to avoid deadlocks
	close(channel)
}

func calculation(base int, channel chan int) {
	time.Sleep(time.Second)
	channel <- 2 * base

	time.Sleep(time.Second)
	channel <- 3 * base

	time.Sleep(3 * time.Second)
	channel <- 4 * base
}

func main() {
	//Create a channel of int with buffer 1
	channel := make(chan int, 1)

	channel <- 1 //send data to channel (write)
	<-channel    //receiving data from channel (read)

	channel <- 2
	fmt.Println(<-channel)

	ch := make(chan int)
	go calculation(2, ch)

	//Wait data to be ready
	a, b := <-ch, <-ch //receiving data from channel
	fmt.Println(a, b)

	fmt.Println(<-ch)

	//Using range and close in channels
	ch2 := make(chan int, 30)
	go primes(cap(ch2), ch2)

	for i := range ch2 {
		fmt.Printf("%d ", i)
	}

	fmt.Println("End")
}
