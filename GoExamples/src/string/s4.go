package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// 易错点：Go 里 len(s) 是得到 s 的 bytes number；
	// utf8.RuneCountInString(s) 才是 string 的长度
	word1 := "Señor"
	fmt.Printf("String: %s\n", word1)
	fmt.Printf("Length of the string: %d\n", utf8.RuneCountInString(word1))
	fmt.Printf("number of bytes: %d\n", len(word1))

	fmt.Println("=======================")
	word2 := "Cat"
	fmt.Printf("String: %s\n", word2)
	fmt.Printf("Length of the string: %d\n", utf8.RuneCountInString(word2))
	fmt.Printf("number of bytes: %d\n", len(word2))
}
