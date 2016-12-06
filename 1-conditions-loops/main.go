package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var (
		randomNumber      = rand.Intn(100)
		beautifulWeather  bool
		holbertonFounders []string
		school            = "Holberton School"
	)

	if randomNumber > 50 {
		fmt.Printf("my random number is %d and is greater than 50\n", randomNumber)
	} else {
		fmt.Printf("my random number is %d and is less than 50\n", randomNumber)
	}
	fmt.Printf("I am a student of the %s\n", school)

	if test1(school) {
		beautifulWeather = true
		fmt.Printf("It's a beautiful weather!\n")
	}

	if beautifulWeather == true {
		holbertonFounders = []string{"Rudy Rigot", "Sylvain Kalache", "Julien Barbier"}
	}

	for i := 0; i < len(holbertonFounders); i++ {
		fmt.Printf("%s is a founder\n", holbertonFounders[i])
	}
}

func test1(s string) bool {
	return s == "Holberton School"

}
