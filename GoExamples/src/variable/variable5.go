package main

import (
	"fmt"
	"math"
)

func main() {
	a, b := 145.6, 89.2
	c := math.Min(a, b)
	fmt.Println("value of c is", c)
}
