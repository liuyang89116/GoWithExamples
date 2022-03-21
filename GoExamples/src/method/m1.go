package main

import "fmt"

type Employee struct {
	name   string
	age    int
	salary int
}

func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %d", e.name, e.salary)
}

func main() {
	e1 := Employee{
		name:   "Amy",
		age:    30,
		salary: 5000,
	}
	e1.displaySalary()
}
