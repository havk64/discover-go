package main

import "fmt"

// Movie is a struct that follows the format to decode json requests to OMDB API
// (http://www.omdbapi.com)
type Movie struct {
	Title      string  `json:"Title"`
	Year       string  `json:"Year"`
	Rated      string  `json:"Rated"`
	IMDBRating float64 `json:"imdbRating,string"`
	IMDBVotes  string  `json:"imdbVotes"`
}

// searchMovies is a struct to decode the query search at OMDB Api
type searchMovies struct {
	Search []struct {
		IMDBID string
	}
}

// Implementation of String method to "satisfy" fmt Stringer interface on
// Movie struct in order to use as default format with "fmt" package
func (movie *Movie) String() string {
	return fmt.Sprintf(
		"The movie : %s was released in %v - the IMDB rating is %v%% "+
			"with %s votes.", movie.Title, movie.Year,
		int(movie.IMDBRating*10), movie.IMDBVotes)
}
