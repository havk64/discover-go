package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
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

	rating, err := strconv.ParseFloat(decode.IMDBRating, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error converting str to float: %v", err)
	}

	fmt.Printf("The movie : %s was released in %s - the IMDB rating is %v%% "+
		"with %s votes.\n", decode.Title, decode.Year,
		rating*10, decode.IMDBVotes)
}
