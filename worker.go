package main

import (
	"net/http"
	"strings"
	"sync"
	"time"
)

type Result struct {
	StatusCode    int
	ContentLength int64
	ResponseTime  time.Duration
	Err           error
	Word          string
}

func sendRequest(url string, word string) Result {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	start := time.Now()
	resp, err := client.Get(url)
	duration := time.Since(start)

	if err != nil {
		return Result{Err: err, Word: word}
	}

	defer resp.Body.Close()

	return Result{
		StatusCode:    resp.StatusCode,
		ContentLength: resp.ContentLength,
		ResponseTime:  duration,
		Err:           nil,
		Word:          word,
	}
}

func StartWorkerPool(threads int, targetPattern string, wordChan <-chan string, resultChan chan<- Result) {
	var wg sync.WaitGroup

	for i := 1; i <= threads; i++ {
		wg.Add(1)
		go runWorker(i, targetPattern, wordChan, resultChan, &wg)
	}

	wg.Wait()
}

func runWorker(id int, targetPattern string, wordChan <-chan string, resultChan chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for word := range wordChan {
		url := strings.ReplaceAll(targetPattern, "FUZZ", word)

		res := sendRequest(url, word)

		resultChan <- res
	}
}
