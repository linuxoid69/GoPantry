package dbase

import (
	"database/sql"
	"time"
)

// Movies --
type Movies struct {
	ImdbID          string          `db:"imdb_id"`
	Title           string          `db:"title"`
	ID              sql.NullInt64   `db:"id"`
	Rating          sql.NullFloat64 `db:"rating"`
	ImdbRating      sql.NullFloat64 `db:"imdb_rating"`
	KinopoiskRating sql.NullFloat64 `db:"kinopoisk_rating"`
	KinopoiskID     sql.NullInt64   `db:"kinopoisk_id"`
	ImageURL        string          `db:"image_url"`
	ReleaseYear     int64           `db:"release_year"`
	CreatedAT       *time.Time      `db:"created_at"`
	UpdatedAT       *time.Time      `db:"updated_at"`
	ISWatched       bool            `db:"is_watched"`
	ISSeries        bool            `db:"is_series"`
}
