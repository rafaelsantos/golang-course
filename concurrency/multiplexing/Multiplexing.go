package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

//origin: read-only channel
func foward(origin <-chan string, destiny chan string) {
	for {
		//When origin channel receives data,
		//the information will be foward to destiny
		destiny <- <-origin
	}
}

//multiplexing : mix messages within channel
func multiplexing(input1, input2 <-chan string) <-chan string {
	channel := make(chan string)
	go foward(input1, channel)
	go foward(input2, channel)

	return channel
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
	//Uses only one channel to handle multi-channel data
	channel := multiplexing(
		title("http://www.codeflow.com.br", "https://www.google.com"),
		title("https://www.uol.com.br", "https://www.stackoverflow.com"),
	)

	log.Println(<-channel)
	log.Println(<-channel)
	log.Println(<-channel)
	log.Println(<-channel)
}
