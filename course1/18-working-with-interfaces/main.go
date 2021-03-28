package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
}
type Cat struct {
}
type Cow struct {
}

func (d Dog) Speak() string {
	return "woof!"
}

func (co Cow) Speak() string {
	return "Moo!"
}

func (c Cat) Speak() string {
	return "Maw!"
}

func main() {

	animals := []Animal{Dog{}, Cat{}, Cow{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}

}
