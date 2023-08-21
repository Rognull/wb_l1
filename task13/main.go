package main

import "fmt"

func main() {
	a:=-1
	b:=1

	fmt.Printf("before a = %d b = %d\n", a, b)

	a, b = b, a

	fmt.Printf("after a = %d b = %d\n", a, b)
}