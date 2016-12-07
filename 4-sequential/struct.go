package main

import "fmt"

type Movie struct {
	Title      string  `json:"Title"`
	Year       string  `json:"Year"`
	Rated      string  `json:"Rated"`
	Released   string  `json:"Released"`
	Runtime    string  `json:"Runtime"`
	Genre      string  `json:"Genre"`
	Director   string  `json:"Director"`
	Writer     string  `json:"Writer"`
	Actors     string  `json:"Actors"`
	Plot       string  `json:"Plot"`
	Language   string  `json:"Language"`
	Country    string  `json:"Country"`
	Awards     string  `json:"Awards"`
	Poster     string  `json:"Poster"`
	Metascore  string  `json:"Metascore"`
	IMDBRating float64 `json:"imdbRating,string"`
	IMDBVotes  string  `json:"imdbVotes"`
	IMDBID     string  `json:"imdbID"`
	Type       string  `json:"Type"`
	Response   string  `json:"Response"`
}

type searchMovies struct {
	Search []struct {
		Title  string
		Year   string
		IMDBID string
	}
}

// Implementation of String method to "satisfy" Stringer interface on Movie struct
func (movie *Movie) String() string {
	return fmt.Sprintf(
		"The movie : %s was released in %v - the IMDB rating is %v%% "+
			"with %s votes.", movie.Title, movie.Year,
		int(movie.IMDBRating*10), movie.IMDBVotes)
}
