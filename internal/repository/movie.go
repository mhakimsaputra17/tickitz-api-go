package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

// MovieRepository handles database operations for movies
type MovieRepository struct {
    db *pgxpool.Pool
}

// NewMovieRepository creates a new MovieRepository instance
func NewMovieRepository(db *pgxpool.Pool) *MovieRepository {
    return &MovieRepository{db: db}
}