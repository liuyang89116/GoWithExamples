package main

import "fmt"

func main() {
	employeeAge := map[string]int{
		"Bob": 26,
		"Amy": 30,
		"Jay": 45,
	}
	fmt.Println("Content of the map...")
	for name, age := range employeeAge {
		fmt.Printf("employeeAge[%s] = %d\n", name, age)
	}
	fmt.Println("Delete Amy...")
	delete(employeeAge, "Amy")
	fmt.Println("map after deletion", employeeAge)
}
