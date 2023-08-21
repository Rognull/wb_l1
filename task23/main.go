package main

import "fmt"

func main() {
    arr := []int{1, 2, 3, 4, 5}
    i := 2
    arr = append(arr[:i], arr[i+1:]...)
    fmt.Println(arr) 

	a := []int{1, 2, 3, 4, 5}

	a = a[:i+copy(a[i:], a[i+1:])]
	fmt.Println(a) 

}


//  https://ueokande.github.io/go-slice-tricks/  наглядный пример для работы с слайсами