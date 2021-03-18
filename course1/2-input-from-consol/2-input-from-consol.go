package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// var s string
	// fmt.Scanln(&s)
	// fmt.Println(s)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your test: ")
	str, _ := reader.ReadString('\n')
	fmt.Println(str)

	fmt.Println("Enter your Number: ")
	str, _ = reader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 6400)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f)
}
