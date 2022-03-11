package main

import "fmt"

func main() {
	s1 := "Go"
	s2 := "is awesome"
	// rst := s1 + " " + s2
	
	rst := fmt.Sprintf("%s %s", s1, s2)
	fmt.Println(rst)
}
