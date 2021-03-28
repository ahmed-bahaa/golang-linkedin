package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	f, err := os.Open("newfile.ext")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(f)
	}

	myError := errors.New("My customized error :p ")
	fmt.Println(myError)

	// Another way

	lessonsAttendence := map[string]bool{
		"ab": true,
		"cd": false,
	}

	attended, ok := lessonsAttendence["ab"]
	if ok {
		fmt.Println("ab attended?", attended)
	} else {
		fmt.Println("No information")
	}

}
