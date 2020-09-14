package main

import (
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

func main() {
	results := make(map[string]string)
	c := make(chan requestResult)
	urls := []string{
		"https://www.naver.com/",
		"https://www.google.com/",
		"https://github.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.daum.com/",
		"https://www.kakao.com/",
		"https://www.instargram.com/",
	}

	for _, url := range urls {
		go hitUrl(url, c)
	}

	for i := 0; i < len(urls); i++ {
		// fmt.Println(<-c)
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitUrl(url string, c chan<- requestResult) {
	resp, err := http.Get(url)
	status := "OK"

	if err != nil || resp.StatusCode <= 400 {
		status = "FAIL"
	}
	// fmt.Println("Chanking")

	c <- requestResult{url: url, status: status}

}
