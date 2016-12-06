package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	fmt.Printf("Hello, we are Holberton School\n"+
		"the date is %s\n"+
		"the year is %d\n"+
		"the month is %s\n"+
		"the day is %d\n", t.String(), t.Year(), t.Month(), t.Day())
}
