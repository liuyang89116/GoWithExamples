package main

import "fmt"

func main() {
	for num, i := 10, 1; num <= 15 && i <= 10; num, i = num+1, i+1 {
		fmt.Printf("%d * %d = %d\n", num, i, num*i)
	}
}
