// Sequential REST client

// Requirements:
// The goal of the "Creating a concurrent RESTClient" section is to create a
// basic command line tool and understand the importance of concurrency.

// Using the client created in the previous task, create a command line tool
// that will sequentialily retreive the for IMDB result for any movie
// (for example the Matrix movie would be a GET to http://www.omdbapi.com/?s=Matrix)

// The movie must be specified in the command line using the movie flag.
// If no movie is specified in the command line, it should get all of the
// Batman movies.

// Measure the execution time of your script.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

// fetchMovies is used to fetch both general movie searches and query for
// individual movies using OMDB API
func fetchMovies(form *url.Values, movies interface{}) {
	uri, err := url.Parse("http://www.omdbapi.com")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing url: %v", err)
	}
	// Set the client and the body form
	client := &http.Client{}
	uri.RawQuery = form.Encode()
	// Buid que http request
	req, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "request error: %v", err)
	}
	// Make que http request
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Request error: %v\n", err)
		os.Exit(1)
	}
	// Decode the json content using specified struct
	if err := json.NewDecoder(res.Body).Decode(&movies); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Close que response body
	defer res.Body.Close()
}

func movieSearch(search *string) {
	start := time.Now()
	// Build the query
	form := &url.Values{
		"s": {*search},
	}

	var movies searchMovies
	// Make the query to the Api
	fetchMovies(form, &movies)
	// Iterate on the result to fetch each individual movies from the API
	for _, v := range movies.Search {
		form = &url.Values{
			"i":    {v.IMDBID},
			"plot": {"short"},
			"r":    {"json"},
		}
		var movie Movie
		fetchMovies(form, &movie)
		fmt.Printf("%v\n", movie.String())
	}
	// Measure the execution time
	defer fmt.Printf("Execution time is %vs\n", time.Since(start).Seconds())
}

func main() {
	moviePtr := flag.String("movie", "Batman", "Search for a Movie using the OMDB Api")
	flag.Parse()

	movieSearch(moviePtr)
}
