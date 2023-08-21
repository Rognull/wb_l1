package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}
	result := make([]int, len(nums))  
	var wg sync.WaitGroup
	wg.Add(len(nums))

	for i, n := range nums {
		go func(i, n int) {
			defer wg.Done()
			result[i] = n * n
		}(i, n)
	}

	wg.Wait()

	fmt.Println(result)
}

