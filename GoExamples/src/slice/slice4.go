package main

import "fmt"

func main() {
	nums := [3]int{100, 200, 300}
	fmt.Println("nums before change", nums)
	nums1 := nums[:]
	nums2 := nums[:]
	nums1[0] = 1
	fmt.Println("nums1 change first elem, nums is now", nums) // [1 200 300]
	nums2[0] = 2
	fmt.Println("nums2 change first elem, nums is now", nums) // [2 200 300]
}
