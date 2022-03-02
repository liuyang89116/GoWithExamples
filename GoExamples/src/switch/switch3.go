package main

import "fmt"

func main() {
	letter := "b"
	fmt.Printf("Letter %s is ", letter)
	switch letter {
	case "a", "e", "i", "o", "u":
		fmt.Printf("a vowel.")
	default:
		fmt.Printf("not a vowel.")
	}
}
