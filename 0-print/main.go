// Go Fundamentals = Print format

// The goal of this task is to understand how to use the go print command.

// print to the sdtout Hello, we are Holberton School
// print the current date/time
// print only the year
// print only the month
// print the only the day
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
