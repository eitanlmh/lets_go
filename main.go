package main

import "fmt"

func main() {
	for i := range 100 {

		switch {
		case i%7 == 0 && i%6 == 0:
			fmt.Println("Six Seveeen!")
		case i%6 == 0:
			fmt.Println("six")
		case i%7 == 0:
			fmt.Println("seven")
		default:
			fmt.Println(i)
		}
	}
}
