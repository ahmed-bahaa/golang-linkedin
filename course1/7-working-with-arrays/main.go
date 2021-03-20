package main

import (
	"fmt"
)

func main() {

	var myArr [3]int
	myArr[0] = 10
	myArr[1] = 11
	myArr[2] = 12

	fmt.Println("my array: ", myArr)
	fmt.Println("my array length : ", len(myArr))

	var myStringarr = [2]string{"hi", "man"}

	fmt.Println("my string array: ", myStringarr)
	fmt.Println("my first string array element : ", myStringarr[0])
	fmt.Println("my string array length : ", len(myStringarr))

}
