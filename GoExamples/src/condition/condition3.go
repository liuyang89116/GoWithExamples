package main

import "fmt"

func main() {
	//if num := 23; num%2 == 0 {
	//	fmt.Println("This number", num, "is even.")
	//} else {
	//	fmt.Println("This number", num, "is odd.")
	//}

	// in Go's philosophy, better to return early
	num := 23
	if num%2 == 0 {
		fmt.Println("This number", num, "is even.")
		return
	}

	fmt.Println("This number", num, "is odd.")
}
