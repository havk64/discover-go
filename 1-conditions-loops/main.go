// ===-----------------------------------------------------------------------===
// Go fundamentals - if/else and loops

// Requirements:
// - Generate a pseudo-random integer between 0 and 100 and store it in a
// variable called randomNumber

// - Create an if/else condition to check whether your number is greater or not
// than 50

// - Store in the variable school the string Holberton School
// - Create an if/else condition to check whether the string school is equal to
//  Holberton School

// - Store in the variable beautifulWeather the boolean value true

// - Create an if/else condition to check whether the variable beautifulWeather
// is true or not

// - Store in the variable holbertonFounders the following values Rudy Rigot,
// Sylvain Kalache, Julien Barbier

// - Create a for loop that will iterate over the holbertonFounders
// ===-----------------------------------------------------------------------===
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
