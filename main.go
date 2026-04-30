package main

import (
	"fmt"
	"net/http"
	"time"
)

type Result struct {
	StatusCode 		int
	ContentLength 	int64
	ResponseTime 	time.Duration
	Err 			error	
}

func sendRequest(url string) Result {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	start := time.Now()
	resp, err := client.Get(url)
	duration := time.Since(start)

	if err != nil {
		return Result{Err: err}
	}

	defer resp.Body.Close()

	return Result{
		StatusCode: 		resp.StatusCode,
		ContentLength: 	resp.ContentLength,
		ResponseTime: 	duration,
		Err: 				nil,
	}
}

func main() {
	targetURL := "https://www.google.com"

	fmt.Println("Sending request to:", targetURL)
	fmt.Println("-------------------------------------------")
	res := sendRequest(targetURL)

	if res.Err != nil {
		fmt.Printf("Error occurred: %v\n", res.Err)
	}
	
	fmt.Printf("Status Code: %d | Length %d | Time %v\n", res.StatusCode, res.ContentLength, res.ResponseTime)
}