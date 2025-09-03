package server

type Movie struct {
	Genres      []string `json:"genres"`
    Title       string   `json:"title"`
	ReleaseYear int      `json:"release_year"`
	Rating      int      `json:"rating"`
	IsSeries    bool     `json:"is_series"`
	IsWatched   bool     `json:"is_watched"`
}
