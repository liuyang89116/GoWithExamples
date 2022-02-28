package main

import "fmt"

func main() {
	num := 85
	if num < 50 {
		fmt.Println(num, "is less than 50")
	} else if num > 90 {
		fmt.Println(num, "is greater than 90")
	} else {
		fmt.Println(num, "is greater than 50 and less than 90")
	}
}
