package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://www.google.com/"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("response Type: %T \n", resp)

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	content := string(bytes)
	fmt.Println(content)

}
