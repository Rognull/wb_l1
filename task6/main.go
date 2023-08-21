package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var (
	stopChan  = make(chan struct{})
	stopFlag  = false
	ctx       context.Context
	cancelCtx context.CancelFunc
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go stopWithChannel(&wg) 

	time.Sleep(3 * time.Second)
	fmt.Println("Stop with chanel")
	stopChan <- struct{}{}

	wg.Add(1)
	go stopWithFlag(&wg)

	time.Sleep(2 * time.Second)
	fmt.Println("Stop with global flag")
	stopFlag = true

	wg.Add(1)
	ctx, cancelCtx = context.WithCancel(context.Background())  // так же можно использовать другие контексты, такие как TODO , withvalue, with deadline, withtimetimeout
	go stopWithContext(&wg)

	time.Sleep(3 * time.Second)
	fmt.Println("Stop with context")
	cancelCtx()

	wg.Wait()
	fmt.Println("Etix")
}


func stopWithChannel(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-stopChan:
			fmt.Println("Gorutine stoped with chanel signal")
			return
		default:
			fmt.Println("Gorutine with chanelstop is working")
			time.Sleep(time.Second)
		}
	}
}


func stopWithFlag(wg *sync.WaitGroup) {
	defer wg.Done()
	for !stopFlag {
		fmt.Println("Gorutine with flag stop is working")
		time.Sleep(time.Second)
	}
	fmt.Println("Gorutine stoped with flag signal")
}

 
func stopWithContext(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Gorutine stoped with context")
			return
		default:
			fmt.Println("Gorutine with contextstop is working")
			time.Sleep(time.Second)
		}
	}
}