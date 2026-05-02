package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	URL      string
	Wordlist string
	Threads  int
	Method   string
	Filters  FilterOptions
}


func ParseFlags() *Config {
	var (
		u  = flag.String("u", "", "URL mục tiêu (phải chứa từ khóa FUZZ)")
		w  = flag.String("w", "", "Đường dẫn đến file wordlist")
		t  = flag.Int("t", 10, "Số lượng luồng (threads) chạy song song")
		x  = flag.String("X", "GET", "Phương thức HTTP (GET, POST, PUT, ...)")
		mc = flag.String("mc", "200,204,301,302,307,401,403", "Match HTTP status codes, phân cách bằng dấu phẩy")
		fc = flag.String("fc", "", "Filter HTTP status codes, phân cách bằng dấu phẩy")
	)

	flag.Parse()

	if *u == "" || *w == "" {
		fmt.Println("Lỗi: Bắt buộc phải cung cấp URL (-u) và Wordlist (-w)")
		fmt.Println("Ví dụ: go run . -u https://example.com/FUZZ -w wordlist.txt")
		os.Exit(1)
	}

	if !strings.Contains(*u, "FUZZ") {
		fmt.Println("Lỗi: URL phải chứa từ khóa 'FUZZ'")
		os.Exit(1)
	}

	return &Config{
		URL:      *u,
		Wordlist: *w,
		Threads:  *t,
		Method:   strings.ToUpper(*x),
		Filters: FilterOptions{
			MatchCodes:  parseIntSlice(*mc),
			FilterCodes: parseIntSlice(*fc),
		},
	}
}

func parseIntSlice(input string) []int {
	if input == "" {
		return nil
	}
	var result []int
	parts := strings.Split(input, ",")
	for _, p := range parts {
		val, err := strconv.Atoi(strings.TrimSpace(p))
		if err == nil {
			result = append(result, val)
		}
	}
	return result
}