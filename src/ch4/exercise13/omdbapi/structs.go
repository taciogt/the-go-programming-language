package omdbapi

type Movie struct {
	Title  string
	Poster string
}

type SearchMovieResponse struct {
	Movies []Movie `json:"Search"`
}
