// Go fundamentals - Struct and Method on Struct

// Requirements:
//Create a user struct that has the following properties:

// - variable Name of type string that exposed to json is name
// - variable DOB of type string that exposed to json is dateofbirth
// - variable City of type string that exposed to json is city
// - initialize the struct user in a variable called u with the following values:
//
// - Betty Holberton as the name
// - March 7, 1917 as the DOB
// - Philadelphia as the city
// - Create a method on the struct that will print Hello
//
// - Create a method on the struct that will produce the following input. who
// was born in would be XX years old today

package main

import (
	"fmt"
	"os"
	"time"
)

type user struct {
	Name string `json: "name"`
	DOB  string `json: "date_of_birth"`
	City string `json: "city"`
}

func main() {

	u := user{
		Name: "Betty Holberton",
		DOB:  "March 7, 1917",
		City: "Philadelphia",
	}
	u.PrintHello()
	if err := u.PrintCity(); err != nil {
		fmt.Fprintf(os.Stderr, "Struct error: %v\n", err)
	}
}

func (u user) PrintHello() {
	fmt.Printf("Hello %s\n", u.Name)

}

func (u user) PrintCity() error {
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
