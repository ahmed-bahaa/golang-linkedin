package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	content := "Hello from Go!"
	file, err := os.Create("./myfile.txt")
	checkError(err)

	ln, err := io.WriteString(file, content)
	checkError(err)

	fmt.Printf("Everything working perfect with %v characters!", ln)

	bytes := []byte(content)
	ioutil.WriteFile("./filefrombytes.txt", bytes, 0644)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
