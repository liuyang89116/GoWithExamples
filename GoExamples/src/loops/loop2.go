package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i > 6 {
			break
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println("\nend for loop")
}
