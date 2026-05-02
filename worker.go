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

func sendRequest(method string, url string, word string) Result {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return Result{Err: err, Word: word}
	}

	start := time.Now()
	resp, err := client.Do(req)
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

func StartWorkerPool(cfg *Config, wordChan <-chan string, resultChan chan<- Result) {
	var wg sync.WaitGroup

	for i := 1; i <= cfg.Threads; i++ {
		wg.Add(1)
		go runWorker(cfg, wordChan, resultChan, &wg)
	}

	wg.Wait()
}

func runWorker(cfg *Config, wordChan <-chan string, resultChan chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for word := range wordChan {
		url := strings.ReplaceAll(cfg.URL, "FUZZ", word)
		res := sendRequest(cfg.Method, url, word)

		if cfg.Filters.IsValid(res) {
			resultChan <- res
		}		
	}
}
