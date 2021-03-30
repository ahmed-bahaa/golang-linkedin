package main

import (
	"fmt"
)

func main() {

	x := 12.123

	// control the decimal precision
	fmt.Printf("%.2f\n", x)

	// control the full length , default is 10
	fmt.Printf("%10f\n", x)

	// print with padding & precision
	fmt.Printf("%10.2f\n", x)

	// print with the sign
	fmt.Printf("%+10.2f\n", x)

	// pad with zeros instead of spaces
	fmt.Printf("%010.2f\n", x)
}
