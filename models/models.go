package models

import (
	"database/sql"
	"time"
)

// Models is the wrapper for the database models
type Models struct{
	DB DBModel
}

// NewModels returns Models with db pool
func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{
			DB: db,
		},
	}
}

// Movie is the type for the movie model
type Movie struct {
	ID          int  `json:"id"`
	Title			 string `json:"title"`
	Description string `json:"description"`
	Year 			int    `json:"year"`
	ReleaseDate time.Time `json:"release_date"`
	Runtime 	 int `json:"runtime"`
	Rating 		 float64 `json:"rating"`
	MPAARating string `json:"mpaa_rating"`
	CreatedAt   time.Time `json:"-"` 
	UpdatedAt   time.Time `json:"-"`
	MovieGenres map[int]string `json:"genres"`
}
 
// Genre is the type for the genre model
type Genre struct {
	ID 				int  `json:"id"`
	GenreName string `json:"genre_name"`
	CreatedAt   time.Time `json:"-"` 
	UpdatedAt   time.Time `json:"-"`
}

// MovieGenre is the type for the movie_genre model
type MovieGenre struct {
	ID 				int  `json:"-"`
	MovieID int `json:"-"`
	GenreID int `json:"-"`
	Genre Genre `json:"genre"`
	CreatedAt   time.Time `json:"-"` 
	UpdatedAt   time.Time `json:"-"`
}