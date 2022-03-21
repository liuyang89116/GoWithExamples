package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	length int
	width  int
}

type Circle struct {
	r float64
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	rect := Rectangle{
		length: 5,
		width:  6,
	}
	fmt.Println("Area of rectangle:", rect.Area())

	circle := Circle{
		r: 6,
	}
	fmt.Println("Area of circle is", circle.Area())
}
