package tv

type TVResponse struct {
	Page       int  `json:"page"`
	TotalPages int  `json:"total_pages"`
	Results    []TV `json:"results"`
}

type TV struct {
	Name         string  `json:"name"`
	FirstAirDate string  `json:"first_air_date"`
	VoteAverage  float32 `json:"vote_average"`
	PosterPath   string  `json:"poster_path"`
	BackdropPath string  `json:"backdrop_path"`
}
