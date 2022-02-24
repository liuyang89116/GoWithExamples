package main

import "fmt"

func main() {
	const a = 5
	var intVar int = a
	var int32Var int32 = a
	var float64Var float64 = a
	var complex64Var complex64 = a

	fmt.Println("intVar:", intVar)
	fmt.Println("int32Var:", int32Var)
	fmt.Println("float64Var:", float64Var)
	fmt.Println("complex64Var:", complex64Var)

	var num = 9.8 / 8
	fmt.Printf("Type of num is %T, value is %v", num, num)
}
