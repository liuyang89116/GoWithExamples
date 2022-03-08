package main

import "fmt"

func main() {
	employeeAge := map[string]int{
		"Bob": 26,
		"Amy": 30,
		"Jay": 45,
	}
	employee := "joe"
	value, flag := employeeAge[employee]
	if flag == true {
		fmt.Println("age of employee is", value)
		return
	}
	fmt.Println(employee, "not found")
}
