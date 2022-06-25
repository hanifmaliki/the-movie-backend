package get

var api_key = "611c4ae739972b9bb9cb9a96393fe3b3"

type TopRated struct {
	Page         int               `json:"page"`
	TotalPages   int               `json:"total_pages"`
	TotalResults int               `json:"total_results"`
	Results      []TopRatedResults `json:"results"`
}

type TopRatedResults struct {
	Adult         bool   `json:"adult"`
	BackdropPath  string `json:"backdrop_path"`
	OriginalTitle string `json:"original_title"`
}
