package main

import (
	"fmt"
	"strings"
)

func reverseWords(data string) string {
	words := strings.Fields(data)
	reversedWords := make([]string, len(words))

	for i, word := range words {
		reversedWords[len(words)-1-i] = word
	}

	return strings.Join(reversedWords, " ")
}

func main() {
	data := "snow dog sun"
	result := reverseWords(data)
	fmt.Println(result)
}