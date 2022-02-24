package main

import "fmt"

func main() {
	a, b := 5.68, 8.59
	fmt.Printf("type of a: %T, b: %T\n", a, b)
	sum := a + b
	diff := a - b
	fmt.Println("sum:", sum, ", diff:", diff)

	num1, num2 := 20, 10
	fmt.Println("sum:", num1+num2, "diff:", num1-num2)
}
