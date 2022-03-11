package main

import "fmt"

func mutate(s []rune) string {
	s[0] = 'a'
	return string(s)
}

func main() {
	s := "hello"
	fmt.Println(mutate([]rune(s))) // string 是 immutable 的，只有变成 a slice of rune 才能 change
}
