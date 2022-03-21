package main

import "fmt"

type employee struct {
	name string
	age  int
}

func (e employee) changeName(name string) {
	e.name = name
}

func (e *employee) changeAge(age int) {
	e.age = age
}

func main() {
	e := employee{
		name: "Bob",
		age:  26,
	}
	fmt.Printf("before change, employee of name %s, age %d\n", e.name, e.age)

	e.changeName("Jim")
	fmt.Printf("change name, employee of name %s, age %d\n", e.name, e.age) // name Bob, age 18

	e.changeAge(18)
	fmt.Printf("change age, employee of name %s, age %d\n", e.name, e.age) // name Bob, age 26
}
