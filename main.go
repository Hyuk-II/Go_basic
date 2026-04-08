package main

import (
	"fmt"
	"net/http"
)

type requestResult struct {
	url string
	status string
}

// 데이터를 보낼 수 만 있는 채널
func hitURL(url string, c chan<- requestResult) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c<-requestResult{url:url, status: status}
}

func main() {
	c := make(chan requestResult)
	results := map[string]string{}
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
		}

	for _, url := range urls {
		go hitURL(url, c)
	}

	fmt.Println("Waiting messages ... ")

	// 여기서 출력하면 받대로 좌랄락
	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}
