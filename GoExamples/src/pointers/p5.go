package main

import "fmt"

func change(val *int) {
	*val = 55
}

func swap(a *int, b *int) {
	tmp := *a;
	*a = *b
	*b = tmp
}

func main() {
	a := 58
	fmt.Println("value of a before call is", a)
	change(&a)
	fmt.Println("value of a after call is", a)

	fmt.Println("=================")
	num1 := 1
	num2 := 2
	swap(&num1, &num2)
	fmt.Printf("value of num1: %d, num2: %d\n", num1, num2)
}
