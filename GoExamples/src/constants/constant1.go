package main

import "fmt"

func main() {
	const a = 50
	fmt.Println(a)
	const (
		name    = "Bob"
		age     = 29
		country = "USA"
	)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(country)
}
