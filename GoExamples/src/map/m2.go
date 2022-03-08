package main

import "fmt"

func main() {
	employeeAge := map[string]int{
		"Bob": 26,
		"Amy": 30,
		"Jay": 45,
	}
	fmt.Println("employee ages", employeeAge)

	age := employeeAge["Bob"]
	fmt.Println("Bob is", age, "yeas old")
}
