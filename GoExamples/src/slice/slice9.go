package main

import "fmt"

func main() {
	arr1 := []string{"a", "b", "c"}
	arr2 := []string{"d", "e", "f"}
	arr := append(arr1, arr2...)
	fmt.Println("arr is", arr)
}
