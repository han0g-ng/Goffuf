package main

import (
	"bufio"
	"os"
	"strings"
)

func ReadWordlist(path string) (<-chan string, error) {
	wordChan := make(chan string)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	go func() {
		defer file.Close()
		defer close(wordChan)

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line != "" {
				wordChan <- line
			}
		}
	}()

	return wordChan, nil
}