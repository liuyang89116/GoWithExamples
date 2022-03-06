package main

import "fmt"

func main() {
	arr := [][]string{
		{"a", "b"},
		{"c", "d"},
		{"e", "f"},
	}

	for _, v1 := range arr {
		for _, v2 := range v1 {
			fmt.Printf("%s", v2)
		}
		fmt.Printf("\n")
	}
}
