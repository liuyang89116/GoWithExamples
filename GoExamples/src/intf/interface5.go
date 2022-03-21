package main

import "fmt"

func assert(i interface{}) {
	v, ok := i.(int)
	fmt.Println(v, ok)
}

func main() {
	var s interface{} = 56
	assert(s) // 56 true
	var name interface{} = "John"
	assert(name) // 0 false
}
