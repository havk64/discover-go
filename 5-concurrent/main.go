package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"
)

// fetchMovies is used to fetch both general movie searches and query for
// individual movies using OMDB API
func fetchMovies(form *url.Values, movies interface{}) <-chan bool {
	ch := make(chan bool)
	go func() {
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
		ch <- true
		close(ch)
	}()
	return ch
}

func movieSearch(search *string) {
	var wg sync.WaitGroup
	start := time.Now()
	// Build the query
	form := &url.Values{
		"s": {*search},
	}

	var movies searchMovies
	// Make the query to the Api
	<-fetchMovies(form, &movies)
	// Initialize the collection to get the list of movies
	result := make([]*Movie, len(movies.Search))
	// Iterate on the result to fetch each individual movies from the API
	for i, v := range movies.Search {
		wg.Add(1)
		go func(i int, v string, result []*Movie) {
			form = &url.Values{
				"i":    {v},
				"plot": {"short"},
				"r":    {"json"},
			}
			var movie Movie
			<-fetchMovies(form, &movie)
			result[i] = &movie
			defer wg.Done()
		}(i, v.IMDBID, result)
	}
	// Wait for the async calls to return
	wg.Wait()
	// Iterate through the final list to print the result
	for _, m := range result {
		fmt.Printf("%v\n", m.String())
	}
	// Measure the execution time
	defer fmt.Printf("Execution time is %v\n", time.Since(start))
}

func main() {
	// Using the flag package to parse the command line arguments
	moviePtr := flag.String("movie", "Batman", "Search for a Movie using the OMDB Api")
	flag.Parse()

	movieSearch(moviePtr)
}
