package main

import "fmt"

type Employee1 struct {
	firstName string
	lastName  string
	age       int
	salary    int
}

func main() {
	emp1 := &Employee1{
		firstName: "Amy",
		lastName:  "Zhang",
		age:       30,
		salary:    1000,
	}

	fmt.Println("first name:", (*emp1).firstName)
	fmt.Println("last name:", emp1.lastName) // go gives the option to omit *
}
