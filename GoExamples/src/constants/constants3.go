package main

import "fmt"

func main() {
	// var defaultName = "Sam"
	type myString string
	var customName myString = "Sam" // allow
	//customName = defaultName // not allowed
	fmt.Println(customName)

	const trueConst = true
	type myBool bool
	var defaultBool = trueConst
	var customBool myBool = trueConst
	fmt.Println(defaultBool)
	fmt.Println(customBool)
	//defaultBool = customBool // not allowed
}
