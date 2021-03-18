package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {

	i1, i2, i3 := 5, 6, 7
	intsum := i1 + i2 + i3
	fmt.Println("Integer sum: ", intsum)

	f1, f2, f3 := 5.2, 6.6, 7.1
	floatsum := f1 + f2 + f3
	fmt.Println("Integer sum: ", floatsum)

	var b1, b2, b3, bigSum big.Float
	b1.SetFloat64(5.2)
	b2.SetFloat64(6.6)
	b3.SetFloat64(7.1)
	bigSum.Add(&b1, &b2).Add(&bigSum, &b3)

	fmt.Printf("Bigsum = %0.1g\n", &bigSum)

	circleRaduis := 6.5
	circumference := circleRaduis * math.Pi
	fmt.Println("circumference : %.2f", circumference)

}
