package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	 
	go Write(ch)

	go Read(ch)
 
	time.Sleep(time.Duration(10) * time.Second)

	close(ch)

	fmt.Println("Exit.")
}

func Write(ch chan<- int) {
	defer close(ch)
	for i := 1; ; i++ {
		fmt.Println("Write value in chanel:", i)
		ch <- i
		time.Sleep(time.Second)
	}
}

func Read(ch <-chan int) {
	for value := range ch {
		fmt.Println("REad value in chanel:", value)
	}
}