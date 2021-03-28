package main

import (
	"fmt"
)

func main() {

	// defer work as stack LIFO
	// it is usefull for example in case handling connections with db to close connection after
	// anything is done
	// or we can use it to make sure to clean after we finished all the steps to cleanup
	defer fmt.Println("statment 1")
	defer fmt.Println("statment 2")
	defer fmt.Println("statment 3")

	doSomething()

	defer fmt.Println("statment 4")
	defer fmt.Println("statment 5")

	fmt.Println("undeferented statment!")

	// things are evaluated first then stored in stack
	x := 1000
	defer fmt.Println("X value: ", x)
	x++
	fmt.Println("the x value after increment :", x)
}

func doSomething() {

	// defer scope is within the function not the whole program
	defer fmt.Println("defered func inside doSomething function")
	fmt.Println("Hello from Something!")

}
