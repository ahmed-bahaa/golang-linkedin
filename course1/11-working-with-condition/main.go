package main

import (
	"fmt"
)

func main() {

	x := 20
	var result string

	if x < 0 {
		result = " value is lower than zero"
	} else if x == 0 {
		result = " value is equal to zero"
	} else {
		result = " value is greater than zero"
	}

	fmt.Println("Result is : ", result)
	fmt.Println(x)

	// This way if we want to define local variable and GC will clean it after the if statment block

	if y := -4; y < 0 {
		result = " value is lower than zero"
	} else if y == 0 {
		result = " value is equal to zero"
	} else {
		result = " value is greater than zero"
	}

	fmt.Println("Result is : ", result)
	// fmt.Println(y)

}
