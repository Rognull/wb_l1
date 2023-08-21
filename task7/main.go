package main

import (
	"fmt"
	"sync"
)

type MapWithMutex struct {
	mu sync.Mutex
	data map[string]int
}

func (m *MapWithMutex)Add (key string, value int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value
}

func (m *MapWithMutex) Get(key string) int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.data[key]
}

func main() {
	mapWithMutex := MapWithMutex{data: make(map[string]int)}

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", i)
			mapWithMutex.Add(key, i)
			fmt.Printf("Wrote %d to %s\n", i, key)
		}(i)
	}

	wg.Wait()

	fmt.Println(mapWithMutex.data)

	// в пакете sync есть структура sync.Map обеспечивающая безопасную работу с мапами

	var m sync.Map

	m.Store("key1", 1)
	m.Store("key2", 2)

	val, ok := m.Load("key1")
	if ok {
		fmt.Println("Value for key1:", val)
	} else {
		fmt.Println("Not found in map")
	}

	m.Delete("key1")

	val, ok = m.Load("key1")
	if ok {
		fmt.Println("Value for key1:", val)
	} else {
		fmt.Println("Not found in map")
	}
}
 