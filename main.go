package main

import (
	"errors"
	"fmt"
)

func divide(num_one, num_two float64) (float64, error) {
	if num_two == 0 {
		return num_one, errors.New(("Number two is 0"))
	}
	divided := num_one / num_two
	return divided, nil
}

func main() {
	num, err := divide(5, 0)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("divided number:", num)
	}
	num, err = divide(5, 1)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("divided number:", num)
	}
}
