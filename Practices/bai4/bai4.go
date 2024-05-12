package main

import (
	"fmt"
	"strings"
)

func formatString(text string) string { // giữ lại các kí tự là chữ cái, dấu cách
	Formattext := strings.Map(func(r rune) rune {
		switch {
		case (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == ' ':
			return r
		}
		return -1

	}, text)
	return Formattext
}

func main() {
	text := "&&Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."
	newtext := formatString(text)
	fmt.Println(formatString(newtext))
	// tách từng từ
	subText := strings.Fields(newtext)
	// đếm tổng số lần xuất hiện
	wordCount := make(map[string]int) // dùng để số lần xuất hiện của từng từ
	for _, word := range subText {
		wordCount[word]++ // key là các từ, còn value là số lần xuất hiện
	}
	fmt.Println("Tổng số lần xuất hiện của từng từ: ", wordCount)

	// tổng số từ
	totalword := len(subText)
	fmt.Println("tổng số từ là : ", totalword)
}
