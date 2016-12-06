// REST client

// Requirements:
// 1. In a struct.go file, create a struct that fullfil the JSON generated by
// http://www.omdbapi.com/?i=tt0372784&plot=short&r=json.

// 2. In main.go, create an HTTP client that will connect execute a GET request
// to the address http://www.omdbapi.com/?i=tt0372784&plot=short&r=json.

// 3. Parse the response of your HTTP client into the struct that you created
// in "Creating a struct from JSON"
// Change the URL ( http://www.omdbapi.com/?i=tt0372784&plot=short&r=json) so
// it can be parametrized.
// Retrieve the information and display it as below. (note that the score need
// to be converted to 100)
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func main() {
	uri, err := url.Parse("http://www.omdbapi.com")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing url: %v", err)
	}

	client := &http.Client{}
	form := &url.Values{
		"i":    {"tt0372784"},
		"plot": {"short"},
		"r":    {"json"},
	}
	uri.RawQuery = form.Encode()

	req, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "request error: %v", err)
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Request error: %v\n", err)
		return
	}

	var movie Movie
	if err := json.NewDecoder(res.Body).Decode(&movie); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("The movie : %s was released in %v - the IMDB rating is %v%% "+
		"with %s votes.\n", movie.Title, movie.Year,
		int(movie.IMDBRating*10), movie.IMDBVotes)
}
