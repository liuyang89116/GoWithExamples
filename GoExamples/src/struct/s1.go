package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	age       int
	salary    int
}

func main() {
	emp1 := Employee{ // better: since it is easy to figure out the field
		firstName: "Amy",
		lastName:  "Zhang",
		age:       30,
		salary:    1000,
	}
	emp2 := Employee{"Bob", "Liu", 26, 10000}

	fmt.Println("Employee1:", emp1)
	fmt.Println("Employee2:", emp2)
	fmt.Printf("Employee1's salary: %d\n", emp1.salary)
	emp2.age = 25
	fmt.Printf("Employee2's age: %d\n", emp2.age)
}
