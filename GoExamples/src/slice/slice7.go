package main

import "fmt"

func main() {
	cars := []string{"Ford", "Toyota", "Honda"}
	fmt.Printf("cars is %s, has a length of %d\n", cars, len(cars))
	cars = append(cars, "Ferrari")
	fmt.Printf("cars is now %s, has a length of %d\n", cars, len(cars))
}
