package main

import (
	"fmt"
	"sort"
)

func main() {

	intSlices := []int{3, 4, 5, 6}
	fmt.Println(intSlices)
	fmt.Println("Length: ", len(intSlices))
	fmt.Println("Cap: ", cap(intSlices))

	// remove first element
	fmt.Println("Removing first element: ")
	intSlices = append(intSlices[1:len(intSlices)])
	fmt.Println(intSlices)

	// remove last element
	fmt.Println("Removing last element: ")
	intSlices = append(intSlices[:len(intSlices)-1])
	fmt.Println(intSlices)

	intmakeSlices := make([]int, 4, 6)
	intmakeSlices[0] = 22
	intmakeSlices[1] = 2
	intmakeSlices[2] = 26
	intmakeSlices[3] = 10

	fmt.Println(intmakeSlices)
	fmt.Println(len(intmakeSlices))
	fmt.Println("cap before appending: ", cap(intmakeSlices))

	intmakeSlices = append(intmakeSlices, 555, 555)

	fmt.Println("cap after appending: ", cap(intmakeSlices))

	intmakeSlices = append(intmakeSlices, 555, 6555, 555, 555, 555, 555, 555, 555, 555)

	fmt.Println("cap after second appending: ", cap(intmakeSlices))

	sort.Ints(intmakeSlices)
	fmt.Println(intmakeSlices)

}
