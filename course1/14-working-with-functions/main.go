package main

import (
	"fmt"
)

func main() {

	doSomething()
	sum := addNumber(3, 5)
	fmt.Println("sum: ", sum)

	sum = addAllVal(3, 5, 10, 65, 36)
	fmt.Println("New sum: ", sum)
}

func doSomething() {
	fmt.Println("doing something")
}

func addNumber(val1, val2 int) int {
	return val1 + val2
}

func addAllVal(values ...int) int {
	fmt.Printf("recieved type %T ", values)
	fmt.Println("")
	fmt.Println(values)
	sum := 0
	for i := range values {
		sum += values[i]
	}
	return sum
}
