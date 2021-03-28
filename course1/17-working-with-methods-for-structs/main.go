package main

import (
	"fmt"
)

type Dog struct {
	Breed  string
	Weight int
	sound  string
}

// this will revieve the value of the object
func (d Dog) speak() {
	d.sound = fmt.Sprintf("%v  \n", d.sound)
	fmt.Println(d.sound)
}

// this will recieve a pointer for the object
func (d *Dog) speakThreeTimes() {
	d.sound = fmt.Sprintf("%v %v %v \n", d.sound, d.sound, d.sound)
	fmt.Println(d.sound)
}

func main() {

	poodle := Dog{"Rex", 60, "woof"}
	fmt.Println(poodle)
	poodle.speak()

	poodle.sound = "Arf"
	poodle.speak()
	poodle.speakThreeTimes()
	poodle.speak()

}
