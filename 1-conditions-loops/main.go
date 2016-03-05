package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//p := fmt.Println
	pf := fmt.Printf
	randomNumber := rand.Intn(100)
	school := "Bad School"
	var beautifulWeather bool
	var holbertonFounders []string
	if randomNumber > 50 {
		school = "Holberton School"
		pf("my random number is %d and is greater than 50\n", randomNumber)
	} else {
		pf("my random number is %d and is less than 50\n", randomNumber)
	}

	pf("I am a student of the %s\n", school)

	if test1(school) {
		beautifulWeather = true
		pf("It's a beatiful weather!\n")
	}

	if beautifulWeather == true {
		holbertonFounders = []string{"Rudy Rigot", "Sylvain Kalache", "Julien Barbier"}
	}
	for i := 0; i < len(holbertonFounders); i++ {
		pf("%s is a founder\n", holbertonFounders[i])
	}
}

func test1(s string) bool {
	return s == "Holberton School"

}
