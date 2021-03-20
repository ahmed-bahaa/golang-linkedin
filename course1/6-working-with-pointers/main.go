package main

import "fmt"

func main() {

	var myIntPointer *int

	if myIntPointer != nil {
		fmt.Println("my pointer value is: ", *myIntPointer)
	} else {
		fmt.Println("my pointer value is: Nil ")
	}

	x := 10
	myIntPointer = &x

	if myIntPointer != nil {
		fmt.Println("my pointer value is: ", *myIntPointer)
	} else {
		fmt.Println("my pointer value is: Nil ")
	}

	fl := 52.366
	flPointer := &fl

	if flPointer != nil {
		fmt.Println("my float pointer value is: ", *flPointer)
	} else {
		fmt.Println("my float pointer value is: Nil ")
	}

	*flPointer = *flPointer * 3

	fmt.Println("my float pointer value is: ", *flPointer)
	fmt.Println("my float original value is: ", fl)

}
