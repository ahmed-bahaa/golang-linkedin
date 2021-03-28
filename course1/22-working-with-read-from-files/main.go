package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	filename := "myfile.txt"

	content, err := ioutil.ReadFile(filename)
	checkError(err)

	result := string(content)
	fmt.Println("content of the file :", result)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
