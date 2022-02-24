package main

import "fmt"

func main() {
	const n = "Bob"
	var name = n
	fmt.Printf("Type %T value %v", name, name)
}
