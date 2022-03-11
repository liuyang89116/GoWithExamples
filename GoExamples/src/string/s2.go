package main

import "fmt"

func charsAndBytePostions(s string) {
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}

func main() {
	name := "Señor"
	charsAndBytePostions(name)
	/*
		S starts at byte 0
		e starts at byte 1
		ñ starts at byte 2 // take 2 bytes
		o starts at byte 4
		r starts at byte 5
	*/
}
