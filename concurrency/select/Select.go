package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"
)

func speak(person string) <-chan string {
	channel := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			channel <- fmt.Sprintf("%s speaking: %d", person, i)
		}
	}()

	return channel
}

func multiplexing(input1, input2 <-chan string) <-chan string {
	channel := make(chan string)
	go func() {
		for {
			select {
			case value := <-input1:
				channel <- value
			case value := <-input2:
				channel <- value
			}
		}
	}()

	return channel
}

func faster(url1, url2, url3 string) string {
	channel1 := title(url1)
	channel2 := title(url2)
	channel3 := title(url3)

	//control struct used for concurrency
	select {
	case time1 := <-channel1:
		return time1
	case time2 := <-channel2:
		return time2
	case time3 := <-channel3:
		return time3
	case <-time.After(1 * time.Millisecond):
		return "All losing"
	default:
		return "No answer"
	}
}

func title(urls ...string) <-chan string {
	channel := make(chan string)
	for _, url := range urls {
		//annonymous function
		go func(url string) {
			response, _ := http.Get(url)
			html, _ := ioutil.ReadAll(response.Body)

			regex, _ := regexp.Compile("<title>(.*?)<\\/title>")
			channel <- regex.FindStringSubmatch(string(html))[1]
		}(url)
	}

	return channel
}

func main() {
	champion := faster(
		"http://www.codeflow.com.br",
		"https://www.youtube.com",
		"https://www.uol.com.br",
	)

	log.Println(champion)

	channel := multiplexing(speak("Rafael"), speak("Jonas"))
	log.Println(<-channel)
	log.Println(<-channel)
	log.Println(<-channel)
	log.Println(<-channel)
	log.Println(<-channel)
	log.Println(<-channel)
}
