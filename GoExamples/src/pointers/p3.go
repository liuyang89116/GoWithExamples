package main

import "fmt"

func main() {
	size := new(int)
	fmt.Printf("size value is %d, type is %T, address is %v\n",
		*size, size, size) // size value is 0, type is *int, address is 0xc0000aa058
	*size = 85
	fmt.Printf("size value is now %d, type is %T, address is %v\n",
		*size, size, size) // size value is now 85, type is *int, address is 0xc0000180a8
}
