package main

import "fmt"

func main() {
	a := 25
	var b *int
	if b == nil {
		fmt.Println("initial address b is", b) // <nil>
		b = &a
		fmt.Println("now address b is", b) // 0xc0000180a8
	}
}
