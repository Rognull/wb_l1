package main

import (
	"fmt"
 
)

func reverseString(input string) string {
	line := []rune(input)
	for i, j := 0, len(line)-1; i < len(line)/2; i, j = i+1, j-1 {
		line[i], line[j] = line[j], line[i]
	}
	return string(line)
}

func main() {
	line := "главрыба — это абырвалг"
	line = reverseString(line)
	fmt.Println(line)
}