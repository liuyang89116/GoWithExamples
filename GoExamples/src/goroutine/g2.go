package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func characters() {
	for i := 'a'; i < 'f'; i++ {
		fmt.Printf("%c ", i)
		time.Sleep(150 * time.Millisecond)
	}
}

func main() {
	go numbers()
	go characters()
	time.Sleep(2 * time.Second)
	fmt.Println("main exit")
}
