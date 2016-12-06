package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var (
		randomNumber      = rand.Intn(100)
		beautifulWeather  = true
		holbertonFounders = []string{
			"Rudy Rigot",
			"Sylvain Kalache",
			"Julien Barbier"}
		school = "Holberton School"
	)

	if randomNumber > 50 {
		fmt.Printf("my random number is %d and is greater than 50\n", randomNumber)
	} else {
		fmt.Printf("my random number is %d and is less than 50\n", randomNumber)
	}

	if school == "Holberton School" {
		fmt.Printf("I am a student of the %s\n", school)
	}

	if beautifulWeather {
		fmt.Printf("It's a beautiful weather!\n")
	}

	for i := 0; i < len(holbertonFounders); i++ {
		fmt.Printf("%s is a founder\n", holbertonFounders[i])
	}
}
