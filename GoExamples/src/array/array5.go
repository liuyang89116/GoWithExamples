package main

import "fmt"

func main() {
	a := [...]string{"a", "b", "c"}
	b := a
	b[0] = "newOne"
	fmt.Println("a is", a) // a is [a b c]
	fmt.Println("b is", b) // b is [newOne b c]
}
