package main

import "fmt"

func main() {
	b := 255
	var a *int = &b
	fmt.Printf("type of a is %T\n", a)
	fmt.Println("address of a is", a) // 0xc0000180a8
}
