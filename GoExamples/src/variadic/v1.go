package main

import "fmt"

func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums) // []int
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "is found at index", i, "in", nums)
			found = true
		}
	}

	if !found {
		fmt.Println(num, "is not found in", nums)
	}
	fmt.Printf("===============\n")
}

func main() {
	find(89, 89, 90, 95)
	find(1, 2, 3, 4, 6)
	find(87)
}
