package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("hello from goroutine")
}

func main() {
	go hello()
	time.Sleep(time.Second) // 让 go 协程可以执行完；否则 main 直接执行完就退了。
	fmt.Println("hello from main")
}
