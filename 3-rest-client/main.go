package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	url := "http://www.omdbapi.com/?i=tt0372784&plot=short&r=json"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	var decode Movie
	if err := json.NewDecoder(res.Body).Decode(&decode); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("The movie : %s was released in %v - the IMDB rating is %v%% "+
		"with %s votes.\n", decode.Title, decode.Year,
		int(decode.IMDBRating*10), decode.IMDBVotes)
}
