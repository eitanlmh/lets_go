package main

import (
	"fmt"
	"time"
)

func main() {

	num_channel := make(chan int)
	letters_channel := make(chan string)

	go func(start int, end int) {
		for i := start; i < end; i++ {
			time.Sleep(100 * time.Millisecond)
			num_channel <- i
		}
	}(1, 10)

	go func() {
		letters := "abcdefghijklmnopqrstuvwxyz"
		for _, value := range letters {
			time.Sleep(500 * time.Millisecond)
			letters_channel <- string(value)
		}
	}()

outer:
	for {
		select {
		case msg1 := <-num_channel:
			fmt.Println("message from num channel:", msg1)
		case msg2 := <-letters_channel:
			fmt.Println("message from letters channel:", msg2)
		case <-time.After(3 * time.Second):
			fmt.Println("timeout")
			break outer
		}

	}

}
