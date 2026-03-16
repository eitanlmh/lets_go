package main

import (
	"fmt"
	"time"
)

func printNumbers(num_one int, num_two int){
	for i := num_one; i < num_two; i++ {
		fmt.Print(i, "\n")
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters(){
	letters := "abcdefgijklmnopqrstuvwxyz"
	for _, letter := range letters{
		fmt.Printf("%c\n", letter)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {

	printNumbers(1, 10)

	printLetters()

	go printNumbers(1, 10)

	go printLetters()

	time.Sleep(20 * time.Second)
}