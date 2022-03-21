package main

import "fmt"

func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("This is string and the value is", i.(string))
	case int:
		fmt.Println("This is int and the value is", i.(int))
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	findType("Messi")
	findType(65)
	findType(89.56)
}
