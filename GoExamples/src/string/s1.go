package main

import "fmt"

func printChars(s string) {
	fmt.Println("Characters:")
	runes := []rune(s) // s is converted to a slice of runes
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c", runes[i])
	}
}

func main() {
	name := "Hello World!"
	fmt.Println("String is", name)
	printChars(name)
}
