package main

import "fmt"

func modify(sls []int) {
	sls[0] = 100
}

func main() {
	a := [3]int{89, 90, 91}
	modify(a[:])
	fmt.Println(a)
}
