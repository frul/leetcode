package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup
	wg.Add(2)
	var mu sync.Mutex
	go func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			mu.Lock()
			counter++
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 1000000; i++ {
			mu.Lock()
			counter--
			mu.Unlock()
		}
	}()
	wg.Wait()
	fmt.Println(counter)
}
