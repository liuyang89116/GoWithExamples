package main

import "fmt"

func main() {
	num := 23
	if num%2 == 0 {
		fmt.Println("This number", num, "is even.")
	} else {
		fmt.Println("This number", num, "is odd.")
	}
}
