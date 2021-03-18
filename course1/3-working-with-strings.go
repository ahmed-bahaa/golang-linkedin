package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "Hi there! implicitly"
	fmt.Println(str1)
	fmt.Printf("This var %v and it is type %T ", str1, str1)

	str2 := "Hi there! explicitly"
	fmt.Printf("This var %v and it is type %T ", str2, str2)

	fmt.Println("")

	fmt.Println(strings.ToUpper(str2))
	fmt.Println(strings.Title(str2))

	val1 := "hello"
	val2 := "Hello"
	fmt.Println("equal? ", val1 == val2)
	fmt.Println("equal? ", strings.EqualFold(val1, val2))

	fmt.Println("string contain exp? ", strings.Contains(str1, "exp"))
	fmt.Println("string contain exp? ", strings.Contains(str2, "exp"))

}
