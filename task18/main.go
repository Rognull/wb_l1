package main

import (
	"fmt"
	"sync"
)

type Worker struct {
	value int
	mu    sync.Mutex
}

func (w *Worker) Increace() {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.value++
}

func (w *Worker) GetValue() int {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.value
}

func (w *Worker)  work( iterations int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < iterations; i++ {
		w.Increace()
	}
}

func main() {
	numOfIncreacers := 4
	iterationsPerThread := 30
	counter := &Worker{}
	var wg sync.WaitGroup

	for i := 0; i < numOfIncreacers; i++ {
		wg.Add(1)
		go counter.work(iterationsPerThread, &wg)
	}

	wg.Wait()

	fmt.Println("Final counter value:", counter.GetValue())
}