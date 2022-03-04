package main

import "fmt"

func main() {
	a := [...]float64{45.6, 23.8, 89.6}
	sum := float64(0)
	//for i, v := range a {
	//	fmt.Printf("%d th element of a is %.2f\n", i, v)
	//	sum += v
	//}
	for _, v := range a {
		sum += v
	}
	fmt.Println("sum of all elements of a is", sum)
}
