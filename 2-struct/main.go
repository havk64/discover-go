package main

import (
	"fmt"
	"time"
)

type user struct {
	Name string
	DOB  string `json: "date_of_birth"`
	City string
}

func main() {

	u := user{
		Name: "Betty Holberton",
		DOB:  "March 7, 1917",
		City: "Philadelphia",
	}
	printHello(u)
	printCity(u)
}

func printHello(u user) {
	fmt.Printf("Hello, %s\n", u.Name)

}

func printCity(u user) error {
	parsing, err := time.Parse("January 2, 2006", u.DOB)
	if err != nil {
		fmt.Println(err)
		return err
	}
	age := time.Since(parsing)
	actualAge := int(age.Hours() / 24 / 365)
	fmt.Printf("%s who was born in %s would be %d years old today.\n", u.Name, u.City, actualAge)
	return nil
}
