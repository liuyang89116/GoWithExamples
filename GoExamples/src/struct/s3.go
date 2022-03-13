package main

import "fmt"

type Person struct {
	string
	int
}

func main() {
	p1 := Person{
		string: "Amy",
		int:    20,
	}
	fmt.Println(p1.string)
	fmt.Println(p1.int)
}
