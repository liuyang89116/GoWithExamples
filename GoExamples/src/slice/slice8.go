package main

import "fmt"

func main() {
	var names []string
	if names == nil {
		fmt.Println("names is nil and going to append")
		names = append(names, "Bob", "John", "Lily")
	}
	fmt.Println("names is now", names)
}
