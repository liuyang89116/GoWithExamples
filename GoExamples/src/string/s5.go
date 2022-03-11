package main

import "fmt"

func compareStr(s1 string, s2 string) {
	if s1 == s2 {
		fmt.Printf("%s and %s are equal\n", s1, s2)
		return
	}
	fmt.Printf("%s and %s are not equal\n", s1, s2)
}

func main() {
	s1 := "Go"
	s2 := "Go"
	compareStr(s1, s2)

	fmt.Println("====================")
	s3 := "hello"
	s4 := "world"
	compareStr(s3, s4)
}
