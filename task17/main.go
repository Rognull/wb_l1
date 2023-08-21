package main

import (
	"fmt"
	"sort"
)

func binarySearch(arr []int, target int) int {  
	index := sort.SearchInts(arr, target) // в пакете sort имеется функция бинарного поиска
	if index < len(arr) && arr[index] == target {
		return index
	}
	return -1
}

func newBinarySearch(data []int, target int) int {
	low := 0
	high := len(data) - 1
	for low <= high {
		middle := low + (high-low)/2 // используется вместо (low + high) / 2 , избавляемся от переполнения в случае если  hight-low/2 слишком большие
		if target < data[middle]{
			high = middle - 1
		} else if target > data[middle]{
			low = middle + 1
		} else {
			return middle
		} 
	}
	return -1
}

func main() {
	arr := []int{1,3,4,5,7,8,9}
	target := 3
	index := binarySearch(arr, target)
	if index >= 0 {
		fmt.Printf("number %d at index %d\n", target, index)
	} else {
		fmt.Println("Cant find\n", target)
	}

	index  = newBinarySearch(arr,target)
	if index >= 0 {
		fmt.Printf("number %d at index %d\n", target, index)
	} else {
		fmt.Println("Cant find\n", target)
	}


}