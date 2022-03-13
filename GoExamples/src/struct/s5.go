package main

import "fmt"

type Address1 struct {
	city  string
	state string
}

type Student1 struct {
	name string
	age  int
	Address1
}

func main() {
	p := Student1{
		name: "Amy",
		age:  20,
		Address1: Address1{
			city:  "Seattle",
			state: "WA",
		},
	}

	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("City:", p.city)   // promoted field
	fmt.Println("State:", p.state) // promoted field
}
