package main

import (
	"fmt"
)

type Dog struct {
	Breed  string
	Weight int
}

func main() {

	poodle := Dog{"Rex", 60}
	fmt.Println(poodle)
	fmt.Printf("%+v \n", poodle)
	fmt.Printf("Breed: %v\nWeight: %v", poodle.Breed, poodle.Weight)

}
