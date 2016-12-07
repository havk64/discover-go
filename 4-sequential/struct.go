package main

type Movie struct {
	Title      string  `json:"Title"`
	Year       int     `json:"Year,string"`
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

type search struct {
	Result []struct {
		Title  string
		Year   int    `json:"Year,string"`
		IMDBID string `json:"imdbID"`
	}
}