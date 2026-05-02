package main

import (
	"fmt"
	"strings"
)

func StartPrinter(resultChan <-chan Result) <-chan bool {
	done := make(chan bool)

	go func() {
		fmt.Printf("%-25s | %-8s | %-10s | %-8s\n", "Payload", "Status", "Length", "Time")
		fmt.Println(strings.Repeat("-", 60))

		for res := range resultChan {
			if res.Err != nil {
				continue
			}
			fmt.Printf("%-15s | %-8d | %-10d | %-8s\n", res.Word, 
			 res.StatusCode, res.ContentLength, res.ResponseTime)
		}
		done <- true
	}()
	return done
}