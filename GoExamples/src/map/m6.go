package main

import "fmt"

func main() {
	employeeAge := map[string]int{
		"Bob": 26,
		"Amy": 30,
		"Jay": 45,
	}
	fmt.Println("map length is", len(employeeAge))
	fmt.Println("Before modify", employeeAge) // [Amy:30 Bob:26 Jay:45]
	employeeAge["Jay"] = 50
	fmt.Println("After modify", employeeAge) // [Amy:30 Bob:26 Jay:50]
}
