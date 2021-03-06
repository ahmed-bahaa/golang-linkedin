package main

import (
	"fmt"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello from the Go web server!</h1>")
}

func main() {

	var h Hello
	err := http.ListenAndServe("localhost:4000", h)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
