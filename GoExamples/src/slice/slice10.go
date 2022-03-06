package main

import "fmt"

func subtract(nums []int) {
	for i := range nums {
		nums[i] -= 2
	}
}

func main() {
	nums := []int{6, 7, 8}
	fmt.Println("slice before call is", nums) // [6 7 8]
	subtract(nums)
	fmt.Println("slice after call is", nums) // [4 5 6]
}
