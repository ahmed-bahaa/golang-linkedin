package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Enter your line:")
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	fmt.Println("You entered :", s)

}
