package movies

type MovieResponse struct {
	Page       int     `json:"page"`
	TotalPages int     `json:"total_pages"`
	Results    []Movie `json:"results"`
}

type Movie struct {
	Title        string  `json:"title"`
	ReleaseDate  string  `json:"release_date"`
	VoteAverage  float32 `json:"vote_average"`
	PosterPath   string  `json:"poster_path"`
	BackdropPath string  `json:"backdrop_path"`
}
