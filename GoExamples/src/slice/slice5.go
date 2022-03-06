package main

import "fmt"

func main() {
	arr := [...]string{"apple", "banana", "grape", "pear"}
	sl := arr[0:2]
	fmt.Println("slice is", sl)
	fmt.Println("length of the slice", len(sl))   // 2
	fmt.Println("capacity of the slice", cap(sl)) // 4
	sl = sl[:cap(sl)]
	fmt.Println("re-slice, slice now is", sl) // [apple banana grape pear]
}
