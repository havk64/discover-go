package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func fetchMovies(form *url.Values, movies interface{}) {
	uri, err := url.Parse("http://www.omdbapi.com")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing url: %v", err)
	}

	client := &http.Client{}
	uri.RawQuery = form.Encode()
	req, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "request error: %v", err)
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Request error: %v\n", err)
		os.Exit(1)
	}

	if err := json.NewDecoder(res.Body).Decode(&movies); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer res.Body.Close()
}

func movieSearch(search *string) {
	form := &url.Values{
		"s": {*search},
	}

	var movies searchMovies

	fetchMovies(form, &movies)

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
}

func main() {
	moviePtr := flag.String("movie", "Batman", "Search for a Movie using the OMDB Api")
	flag.Parse()

	movieSearch(moviePtr)
}
