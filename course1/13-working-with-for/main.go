package main

import (
	"fmt"
)

func main() {

	sum := 1
	fmt.Println("Sum: ", sum)

	colors := []string{"red", "blue", "black"}
	fmt.Println("Colors: ", colors)

	for i := 0; i < 10; i++ {
		sum += sum
	}

	fmt.Println("Sum: ", sum)

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	for i := range colors {
		fmt.Println(colors[i])
	}

	sum = 1

	for sum < 1000 {
		sum += sum
		fmt.Println("Sum:", sum)
		if sum > 200 {
			goto endoftheprogram
		}
		if sum > 500 {
			break
		}

	}

endoftheprogram:
	fmt.Println("ENd of the program")
}
