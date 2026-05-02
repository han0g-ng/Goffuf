package main

import (
	"fmt"
	"strings"
)



func main() {
	cfg := ParseFlags()

	wordChan, err := ReadWordlist(cfg.Wordlist)
	if err != nil {
		fmt.Printf("Lỗi đọc wordlist: %v\n", err)
		return
	}

	resultChan := make(chan Result)

	fmt.Printf("\nBắt đầu Fuzzing tới: %s\n", cfg.URL)
	fmt.Printf("Method: %s | Threads: %d | Wordlist: %s\n", cfg.Method, cfg.Threads, cfg.Wordlist)
	fmt.Println(strings.Repeat("-", 60))

	printerDone := StartPrinter(resultChan)

	go func(){
		StartWorkerPool(cfg, wordChan, resultChan)
		close(resultChan)

	}()

	<-printerDone


	fmt.Println(strings.Repeat("-", 60))

	fmt.Println("Hoàn tất")
	
}