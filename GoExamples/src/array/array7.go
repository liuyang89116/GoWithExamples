package main

import "fmt"

func main() {
	a := [...]float64{12.3, 89.5, 45.6}
	fmt.Println("length of the array", len(a))
	for i := 0; i < len(a); i++ {
		fmt.Printf("%dth element of a is %.2f\n", i, a[i])
	}
}
