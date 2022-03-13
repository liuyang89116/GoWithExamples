package main

import "fmt"

type Address struct {
	city  string
	state string
}

type Student struct {
	name    string
	age     int
	address Address
}

func main() {
	p := Student{
		name: "Amy",
		age:  20,
		address: Address{
			city:  "Seattle",
			state: "WA",
		},
	}

	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("City:", p.address.city)
	fmt.Println("State:", p.address.state)
}
