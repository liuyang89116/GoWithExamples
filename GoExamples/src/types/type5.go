package main

import "fmt"

func main() {
	i := 10
	var j float64 = float64(i) // i has to first convert to float
	fmt.Println("j:", j)
}
