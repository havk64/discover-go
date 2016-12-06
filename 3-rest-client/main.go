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
