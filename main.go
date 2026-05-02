package main

import (
	"fmt"
	"strings"
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
	targetURL := "https://www.google.com/FUZZ"
	wordlistPath := "D:\\tools\\ffuf\\fuzz.txt"

	wordChan, err := ReadWordlist(wordlistPath)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Printf("%-15s | %-8s | %-10s | %s\n", "Payload", "Status", "Length", "Time")
    fmt.Println(strings.Repeat("-", 50))

	for word := range wordChan {
		url := strings.ReplaceAll(targetURL, "FUZZ", word)

		res := sendRequest(url)

		if res.Err != nil {
			fmt.Printf("%-15s | ERROR\n", word)
			continue
		}

		fmt.Printf("%-15s | %-8d | %-10d %v\n", word, res.StatusCode, res.ContentLength, res.ResponseTime)
	}
}