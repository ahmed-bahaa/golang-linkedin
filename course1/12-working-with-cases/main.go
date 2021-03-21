package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())
	day := rand.Intn(6) + 1
	fmt.Println(day)

	result := ""

	switch day {
	case 2:
		result = "Monday"
	case 4:
		result = "wednesday"
	default:
		result = "It's a weekday"
	}

	fmt.Println(result)

	//  We can oinitilaize the value to be the local scope
	switch x := -24; {
	case x < 0:
		result = "less than zero"
	case x > 0:
		result = "greater than zero"
	default:
		result = "equal to zero"
	}

	fmt.Println(result)

}
