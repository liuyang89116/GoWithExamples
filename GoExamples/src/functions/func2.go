package main

import "fmt"

//func rectangleCalc(length, width float64) (float64, float64) {
//	var area = length * width
//	var perimeter = (length + width) * 2
//	return area, perimeter
//}

func rectangleCalc(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return
}

func main() {
	area, perimeter := rectangleCalc(10.85, 6.75)
	fmt.Printf("Area is %f, perimeter is %f", area, perimeter)
}
