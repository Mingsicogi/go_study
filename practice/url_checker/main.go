package main

import (
	"fmt"
	"net/http"
)

var urls = []string{
	"https://naver.com",
	"https://google.com",
	"https://daum.net",
	"https://apple.com",
	"https://facebook.com",
	"https://twitter.com",
	"https://line.me",
	"https://blog.naver.com/jieuni4u",
	"https://www.amazon.com",
	"https://www.reddit.com",
}

type result struct {
	url        string
	status     string
	errMessage string
}

func main() {
	channel := make(chan result)
	results := make(map[string]string)
	for _, url := range urls {
		go hitURLForGoroutine(url, channel)
	}

	for i := 0; i < len(urls); i++ {
		channelMessage := <-channel
		results[channelMessage.url] = channelMessage.status
	}

	fmt.Println(results)

	fmt.Println("\n===================================\n")

	for _, url := range urls {
		result, err := hitURL(url)
		if err == nil {
			fmt.Println(result)
		} else {
			fmt.Println(err)
		}
	}
}

func hitURL(url string) (string, error) {
	response, err := http.Get(url)
	if err == nil {
		return url + " hit result : " + response.Status, nil
	} else {
		return "", err
	}
}

// channel send only(<-)
func hitURLForGoroutine(url string, channel chan<- result) {
	response, err := http.Get(url)

	reqResult := result{
		url: url,
	}

	if err != nil {
		reqResult.errMessage = err.Error()
	} else {
		reqResult.status = response.Status
	}

	channel <- reqResult
}

func goRoutine(message string, channel chan string) {
	fmt.Println("receive message", message)
	channel <- "call from go routine function " + message
}
