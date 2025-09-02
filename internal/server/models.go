package server

type Movies []struct {
	Title       string `json:"title"`
	IsSeries    bool   `json:"is_series"`
	Rating      int    `json:"rating"`
	IsWatched   bool   `json:"is_watched"`
	ReleaseYear int    `json:"release_year"`
}
