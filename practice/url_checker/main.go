package main

import (
	"fmt"
	"net/http"
	"strconv"
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
	channel := make(chan string)

	for i := 0; i < 10; i++ {
		go goRoutine(strconv.Itoa(i), channel)
	}

	fmt.Println(<-channel)
	//fmt.Println(<- channel)
	//fmt.Println(<- channel)
	//fmt.Println(<- channel)
	//fmt.Println(<- channel)
	//fmt.Println(<- channel)

	//for _, url := range urls {
	//	err := hitURL(url)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//}

}

func hitURL(url string) error {
	response, err := http.Get(url)
	if err == nil {
		fmt.Println(url, "Request Result :", response.Status)
	}
	return err
}

func goRoutine(message string, channel chan string) {
	fmt.Println("receive message", message)
	channel <- "call from go routine function " + message
}
