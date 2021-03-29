package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
)

type Tour struct {
	Name, Price string
}

func main() {
	url := "http://services.explorecalifornia.org/json/tours.php"
	content := contentFromServer(url)

	fmt.Println(content)
	tours := toursFromJson(content)

	for _, tour := range tours {
		price, _, _ := big.ParseFloat(tour.Price, 10, 2, big.ToZero)
		fmt.Printf("Name: %v and price: %.2f \n", tour.Name, price)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func contentFromServer(url string) string {
	resp, err := http.Get(url)
	checkErr(err)

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	checkErr(err)

	return string(bytes)
}

func toursFromJson(content string) []Tour {
	tours := make([]Tour, 0, 20)

	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	checkErr(err)

	var tour Tour
	for decoder.More() {
		err := decoder.Decode(&tour)
		checkErr(err)
		tours = append(tours, tour)
	}

	return tours
}
