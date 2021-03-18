package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Date(2009, time.November, 11, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at  %s\n", t)

	now := time.Now()
	fmt.Printf("The time is %s\n", now)

	fmt.Println("The Month is ", t.Month())
	fmt.Println("The Day is ", t.Day())
	fmt.Println("The weekDay is ", t.Weekday())

	tomorrow := t.AddDate(0, 0, 1)
	fmt.Printf("Tomorrow is %v, %v %v, %v\n", tomorrow.Weekday(), tomorrow.Month(), tomorrow.Day(), tomorrow.Year())

	longformat := "Monday, January 2, 2006"
	fmt.Println("Tomorrow is ", t.Format(longformat))
	shortformat := "1/2/2006"
	fmt.Println("Tomorrow is ", t.Format(shortformat))
}
