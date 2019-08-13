package main

import (
	"fmt"
	"time"
)

func speak(person, text string, quantity int) {
	for i := 0; i < quantity; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteration %d)\n", person, text, i+1)
	}
}

func main() {
	parallel := true

	if parallel {
		//Go routine only works if main thread is working
		go speak("Maria", "Why not speak to me?", 50)
		go speak("John", "Can only speak after you!", 50)
	} else {
		speak("Maria", "Why not speak to me?", 3)
		speak("John", "Can only speak after you!", 1)
	}
}
