package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func fetchMovies(form *url.Values, movies interface{}) {
	// ch := make(chan bool)
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

func fetchSearch(m string) {
	search := "Batman"
	if m != "" {
		search = m
	}

	form := &url.Values{
		"s": {search},
	}

	var movies searchMovies
	fetchMovies(form, &movies)
	fmt.Printf("%#v\n\n", movies)
	for _, v := range movies.Search {
		// fmt.Printf("%v\n", v.IMDBID)
		form = &url.Values{
			"i":    {v.IMDBID},
			"plot": {"short"},
			"r":    {"json"},
		}
		var movie Movie
		fetchMovies(form, &movie)

		fmt.Printf("The movie : %s was released in %v - the IMDB rating is %v%% "+
			"with %s votes.\n", movie.Title, movie.Year,
			int(movie.IMDBRating*10), movie.IMDBVotes)
	}
}

func main() {
	fetchSearch("Matrix")
}
