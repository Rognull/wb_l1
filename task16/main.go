package main

import (
	"fmt"
)

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)/2]
	var less, equal, greater []int

	for _, num := range arr {
		switch {
		case num < pivot:
			less = append(less, num)
		case num == pivot:
			equal = append(equal, num)
		case num > pivot:
			greater = append(greater, num)
		}
	}

	less = quickSort(less)
	greater = quickSort(greater)

	return append(append(less, equal...), greater...)
}

func main() {
	arr := make([]int,0,10)
	arr = append(arr,1,3,2,4,3,5,4,6,5,7)

	sortedArr := quickSort(arr)
	fmt.Println("Sorted array:", sortedArr)
}