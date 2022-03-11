package main

import "fmt"

func main() {
	//byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}

	// decimal equivalent of {'\x43', '\x61', '\x66', '\xC3', '\xA9'}
	byteSlice := []byte{67, 97, 102, 195, 169}
	s := string(byteSlice)
	fmt.Println(s) // Café

	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	s1 := string(runeSlice)
	fmt.Println(s1) // Señor
}
