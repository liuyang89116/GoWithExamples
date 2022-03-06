package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6}
	sl := arr[1:4]
	fmt.Println("slice is", sl)
	fmt.Println("before arr", arr) // [0 1 2 3 4 5 6]
	for i := range sl {
		sl[i]++
	}
	fmt.Println("after arr", arr) // [0 2 3 4 4 5 6]
}
