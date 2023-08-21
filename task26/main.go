package main

import (
	"fmt"
	"strings"
)

func areUnique(input string) bool {
	unique := make(map[rune]bool)
	for _, char := range input {
		lowerChar := strings.ToLower(string(char))
		if unique[[]rune(lowerChar)[0]] {
			return false
		}
		unique[[]rune(lowerChar)[0]] = true
	}
	return true
}

func main() {
	testStrings := []string{
		"abcd",
		"abCdefAaf",
		"aabcd",
	}

	for _, testStr := range testStrings {
		isUnique := areUnique(testStr)
		fmt.Printf("%s — %t\n", testStr, isUnique)
	}
}