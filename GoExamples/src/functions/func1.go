package main

import "fmt"

func calculateBill(price, cnt int) int {
	var totalPrice = price * cnt
	return totalPrice
}

func main() {
	price, cnt := 85, 6
	totalPrice := calculateBill(price, cnt)
	fmt.Println(totalPrice)
}
