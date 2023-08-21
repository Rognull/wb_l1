package main

import (
	"fmt"
)

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]bool)

	for _, item := range arr {
		set[item] = true
	}

	fmt.Println(set)
 
}