package main

import "fmt"

func main() {
	finger := 5
	fmt.Printf("finger %d is ", finger)
	switch finger {
	case 1:
		fmt.Printf("Thumb")
	case 2:
		fmt.Printf("Index")
	case 3:
		fmt.Printf("Middle")
	case 4:
		fmt.Printf("Ring")
	case 5:
		fmt.Printf("Pinky")
	}
}
