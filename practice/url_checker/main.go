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
}

func main() {
	for _, url := range urls {
		err := hitURL(url)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func hitURL(url string) error {
	response, err := http.Get(url)
	if err == nil {
		fmt.Println(url, "Request Result :", response.Status)
	}
	return err
}
