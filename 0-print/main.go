package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	pf := fmt.Printf
	p := fmt.Println

	p("Hello, we are Holberton School")
	pf("the date is %s\n", t.String())
	//pf("the date is %d-%02d-%02d %d:%d:%d: \n", t.Year(), t.Month(), t.Day(), a, b, c)
	//pf("%s\n", t.String())
	pf("the year is %d\n", t.Year())
	pf("the month is %s\n", t.Month())
	pf("the day is %d\n", t.Day())
}

// 2016-03-05 11:33:01.671605587 -0800 PST
