package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var sum = 0
	for i := 1; i <= 1000; i++ {
		wg.Go(func() {
			mu.Lock()
			defer mu.Unlock()
			sum += i
		})
	}
	wg.Wait()
	fmt.Println("sum:", sum)
}
