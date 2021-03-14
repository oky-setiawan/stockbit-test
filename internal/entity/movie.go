package entity

//GetMovie entity
type (
	GetMovieRequest struct {
		Keyword string
		Page    int
	}

	GetMovieData struct {
		Title  string
		Year   string
		ImdbID string
		Type   string
		Poster string
	}

	GetMovieResponse struct {
		Data []GetMovieData
	}
)

//SearchMovie entity
type (
	SearchMovieRequest struct {
		Keyword string
		Page    int
	}

	SearchMovieData struct {
		Title  string `json:"title"`
		Year   string `json:"year"`
		ImdbID string `json:"imdb_id"`
		Type   string `json:"type"`
		Poster string `json:"poster"`
	}

	SearchMovieResponse struct {
		Data []SearchMovieData `json:"data"`
	}
)
