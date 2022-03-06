package main

import "fmt"

func main() {
	a := [5]int{0, 1, 2, 3, 4}
	b := a[1:4]    // 左开右闭
	fmt.Println(b) // [1 2 3]
}
