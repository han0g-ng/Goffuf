package main

import (
	"fmt"
	"strings"
)



func main() {
	targetPattern := "https://cybershare.league.cyberjutsu-lab.tech/FUZZ"
	wordlistPath := "D:\\tools\\ffuf\\fuzz.txt"
	threads := 10

	filters := FilterOptions{
		MatchCodes:  []int{200, 301, 302, 403, 500},
		FilterCodes: []int{404},
	}

	wordChan, err := ReadWordlist(wordlistPath)

	if err != nil {
		fmt.Printf("Lỗi khởi tạo wordlist: %v\n", err)
		return
	}

	resultChan := make(chan Result)

	fmt.Printf("Bắt đầu Fuzzing với %d workers...\n", threads)
	fmt.Println(strings.Repeat("-", 60))

	printerDone := StartPrinter(resultChan)

	go func(){
		StartWorkerPool(threads, targetPattern, wordChan, resultChan, filters)
		close(resultChan)

	}()

	<-printerDone


	fmt.Println(strings.Repeat("-", 60))

	fmt.Println("Hoàn tất")
	
}