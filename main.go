package main

import (
	"fmt"
	"time"
)

func numbers(numbers_channel chan int, start int, end int) {
	for i := start; i < end; i++ {
		time.Sleep(100 * time.Millisecond)
		numbers_channel <- i
	}
	close(numbers_channel)
}
func letters(letters_channel chan string) {
	letters := "abcdefghijklmnopqrstuvwxyz"
	for _, value := range letters {
		time.Sleep(500 * time.Millisecond)
		letters_channel <- string(value)
	}
	close(letters_channel)
}
func main() {

	num_channel := make(chan int)
	letters_channel := make(chan string)

	go numbers(num_channel, 1, 10)
	go letters(letters_channel)

outer:
	for {
		select {
		case msg1, ok := <-num_channel:
			if ok {
				fmt.Println("message from num channel:", msg1)
			}
		case msg2, ok := <-letters_channel:
			if ok {
				fmt.Println("message from letters channel:", msg2)
			}
		case <-time.After(3 * time.Second):
			fmt.Println("timeout")
			break outer
		}

	}

}
