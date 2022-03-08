package main

import "fmt"

func main() {
	employeeAge := make(map[string]int)
	employeeAge["Bob"] = 26
	employeeAge["Amy"] = 30
	employeeAge["Jay"] = 45
	fmt.Println("employee ages", employeeAge)
}
