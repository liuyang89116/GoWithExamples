package main

import "fmt"

func printArray(arr [3][2]string) {
	for _, v1 := range arr {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

func main() {
	a := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dot"},
		{"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
	}
	fmt.Println("print a")
	printArray(a)

	var b [3][2]string
	b[0][0] = "apple"
	b[0][1] = "pear"
	b[1][0] = "orange"
	b[1][1] = "banana"
	b[2][0] = "tofu"
	b[2][1] = "juice"
	fmt.Println("print b")
	printArray(b)
}
