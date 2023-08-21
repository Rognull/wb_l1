package main

import (
	"fmt"
	"sync"
)

func SumResult (wg *sync.WaitGroup, ch chan int ){
	defer wg.Done()
	 resultSum := 0
		for n := range ch{
			resultSum = resultSum +  n 
		}
		fmt.Println(resultSum)
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	result := make(chan int)
	
 
	var wg sync.WaitGroup
	var wg2 sync.WaitGroup

	wg2.Add(1)
	wg.Add(len(nums))


	go SumResult(&wg2,result)

	for _ , n := range nums {
		go func(n int) {
			defer wg.Done()
			result <- n * n
		}(n)
	}


	wg.Wait()

	close(result)

	wg2.Wait()
	

	 
 
}

