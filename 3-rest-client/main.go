package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	url := "http://www.omdbapi.com/?i=tt0372784&plot=short&r=json"
	req, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("status code is %s\n", req.Status)

	var decoded Movie
	json.NewDecoder(req.Body).Decode(&decoded)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("The movie : %s was released in %s - the IMDB rating is %d with %s votes.\n", decoded.Title, decoded.Year, int(decoded.IMDBRating*10), decoded.IMDBVotes)
}
