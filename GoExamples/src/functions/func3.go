package main

import "fmt"

func rectangleCalc2(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

func main() {
	area, _ := rectangleCalc2(10.8, 6.2)
	fmt.Println("Area is", area)
}
