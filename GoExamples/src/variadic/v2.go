package main

import "fmt"

func change(s ...string) {
	s[0] = "Go"
}

func main() {
	arr := []string{"hello", "world"}
	change(arr...)
	fmt.Println(arr) // [Go world]
}
