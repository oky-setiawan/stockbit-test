package omdb

type omdbResponse struct {
	Response string `json:"Response"`
	Error    string `json:"Error,omitempty"`
}

type omdbGetMovieResponse struct {
	Data  []omdbGetMovieData `json:"Search,omitempty"`
	Total string             `json:"totalResults,omitempty"`
	omdbResponse
}

type omdbGetMovieData struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}
