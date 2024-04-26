package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func Word(s string) map[string]int {
	words := strings.Fields(s) // tách chuỗi s thành từng từ
	wordcnt := make(map[string]int)
	for _, word := range words {
		wordcnt[word]++
	}
	return wordcnt
}
func main() {
	wc.Test(Word)
}
