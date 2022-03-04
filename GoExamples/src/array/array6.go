package main

import "fmt"

func changeLocal(nums [5]int) {
	nums[0] = 60
	fmt.Println("inside function", nums)
}

func main() {
	nums := [...]int{5, 6, 7, 8, 9}
	fmt.Println("before calling function", nums)
	changeLocal(nums)
	fmt.Println("after calling function", nums)
}
