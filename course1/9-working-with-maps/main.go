package main

import (
	"fmt"
	"sort"
)

func main() {

	cities := make(map[string]string)

	fmt.Println(cities)

	cities["CA"] = "California"
	cities["WA"] = "Washington"
	cities["ALX"] = "Alexandria"

	fmt.Println(cities)

	delete(cities, "CA")
	fmt.Println(cities)

	cities["CA"] = "California"
	fmt.Println(cities)

	fmt.Println("UnSorted map: ")

	for k, v := range cities {
		fmt.Printf("%v : %v \n", k, v)
	}

	keys := make([]string, len(cities))
	i := 0
	for k := range cities {
		keys[i] = k
		i++
	}

	sort.Strings(keys)

	fmt.Println("Sorted map: ")

	for k := range keys {
		fmt.Println(cities[keys[k]])
	}

}
