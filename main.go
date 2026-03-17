package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	fmt.Println("worker", id, "started")
	time.Sleep(time.Millisecond * time.Duration(rand.IntN(2000-500)+500))
	fmt.Println("worker", id, "done")
	wg.Done()
}
func main() {
	var wg sync.WaitGroup

	wg.Add(5)
	for i := range 5 {
		worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("all workers finished")
}
