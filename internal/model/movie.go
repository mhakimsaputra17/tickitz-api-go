package model

import "time"

// Movie represents a movie entity from the database
type Movie struct {
    MovieID     int       `json:"movie_id"`
    Title       string    `json:"title"`
    Description string    `json:"description,omitempty"`
    Duration    int       `json:"duration,omitempty"`
    ReleaseDate time.Time `json:"release_date,omitempty"`
    PosterURL   string    `json:"poster_url,omitempty"`
    BackdropURL string    `json:"backdrop_url,omitempty"`
    Rating      float64   `json:"rating,omitempty"`
    CreatedAt   time.Time `json:"created_at,omitempty"`
    UpdatedAt   time.Time `json:"updated_at,omitempty"`
    
    // Add genres field - this will be populated from the Genres table
    Genres      []Genre   `json:"genres,omitempty"`
}

// Genre represents a movie genre from the Genres table
type Genre struct {
    GenreID int    `json:"genre_id"`
    Name    string `json:"name"`
}