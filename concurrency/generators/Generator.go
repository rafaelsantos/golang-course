package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

//Concurrency pattern
//<-chan - read-only channel
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
	title1 := title("http://www.codeflow.com.br", "https://www.google.com")
	title2 := title("https://www.bol.com.br", "https://www.youtube.com")
	fmt.Println(<-title1, <-title2)
	fmt.Println(<-title1, <-title2)
}
