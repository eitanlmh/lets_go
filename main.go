package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (R Rectangle) Area() float64 {
	return R.height * R.width
}

func (C Circle) Area() float64 {
	return math.Pow(C.radius, 2) * math.Pi
}

func totalArea(shapes []Shape) float64 {
	sum := 0.

	for _, shape := range shapes {
		sum += shape.Area()
	}

	return sum
}

func main() {

	shapes := []Shape{Rectangle{2., 3.}, Rectangle{2., 3.}, Circle{5.}, Circle{6.}}
	fmt.Println(totalArea(shapes))
}
