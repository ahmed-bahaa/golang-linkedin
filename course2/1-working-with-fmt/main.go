package main

import "fmt"

type circle struct {
	radius, border int
}

func main() {
	x := 5
	f := 2.6

	// print integer
	fmt.Printf("%d\n", x)
	fmt.Printf("%x\n", x)

	fmt.Printf("%t\n", x > 0)

	// printing a float number
	fmt.Printf("%f\n", f)
	fmt.Printf("%e\n", f)

	// printing using indixing
	fmt.Printf("%[2]d %[1]d\n", 20, 30)
	fmt.Printf("%[1]d %#[1]o %#[1]x\n", 20)

	c := circle{
		radius: 20,
		border: 5,
	}

	// this print the struct values
	fmt.Printf("%v\n", c)

	// this print the struct values and field names
	fmt.Printf("%+v\n", c)

	// this print the struct Type
	fmt.Printf("%T\n", c)

	// sprintf store the strnig ^ u can print it when ever u want
	s := fmt.Sprintf("hello %v  ", 5)
	fmt.Println(s)
}
