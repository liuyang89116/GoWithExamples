package main

import "fmt"

func main() {
	b := 256
	a := &b
	fmt.Println("address of a is", a)
	fmt.Println("value of a is", *a) // 256
	*a++
	fmt.Println("value of a is now", *a) //257
}
