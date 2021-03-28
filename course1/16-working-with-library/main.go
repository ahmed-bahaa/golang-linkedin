package main

import (
	"fmt"
	"stringutil"
)

func main() {
	n1, l1 := stringutil.FullName("ahmed", "bahaa")
	fmt.Printf("Fullname: %v, Number of characters: %v \n\n", n1, l1)

	n2, l2 := stringutil.FullNameNaked("ahmed", "bahaa")
	fmt.Printf("Fullname: %v, Number of characters: %v \n\n", n2, l2)

}
