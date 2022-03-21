package main

import "fmt"

func describe1(i interface{}) {
	fmt.Printf("type %T, value %v\n", i, i)
}

func main() {
	s := "Hello World!"
	describe1(s)
	num := 66
	describe1(num)
	emp := struct {
		name string
	}{
		name: "Messi",
	}
	describe1(emp)
}
