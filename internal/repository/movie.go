package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mhakimsaputra17/tickitz-api-go/internal/model"
)

// MovieRepository handles database operations for movies
type MovieRepository struct {
    db *pgxpool.Pool
}

// NewMovieRepository creates a new MovieRepository instance
func NewMovieRepository(db *pgxpool.Pool) *MovieRepository {
    return &MovieRepository{db: db}
}

// GetMovieGenres fetches all genres for a specific movie
func (r *MovieRepository) GetMovieGenres(movieID int) ([]model.Genre, error){
	query := `
        SELECT g.genre_id, g.name
        FROM genres g
        JOIN movie_genres mg ON g.genre_id = mg.genre_id
        WHERE mg.movie_id = $1
    `

	rows, err := r.db.Query(context.Background(), query, movieID)
	if err!= nil {
		return nil, fmt.Errorf("error querying movie genres: %w",err)
	}
	defer rows.Close()

	var genres []model.Genre
	for rows.Next(){
		var genre model.Genre
		if err := rows.Scan(&genre.GenreID, &genre.Name); err != nil{
			return nil, fmt.Errorf("error scanning genre row: %w", err)
		}
		genres = append(genres, genre)
	}

	if err := rows.Err(); err != nil{
		return nil, fmt.Errorf("error iterating genre rows: %w", err)
	}

	return genres, nil

}

// GetMovies retrieves all movies with pagination and populates their genres
func (r *MovieRepository) GetMovies(page, limit int) ([]model.Movie, int, error) {
    offset := (page - 1) * limit

    // Get total count for pagination
    var totalCount int
    countErr := r.db.QueryRow(context.Background(), "SELECT COUNT(*) FROM movies").Scan(&totalCount)
    if countErr != nil {
        return nil, 0, fmt.Errorf("error counting movies: %w", countErr)
    }

    // Query movies with pagination
    query := `
        SELECT movie_id, title, COALESCE(description, ''), COALESCE(duration, 0), 
               COALESCE(release_date, CURRENT_DATE), COALESCE(poster_url, ''), 
               COALESCE(backdrop_url, ''), COALESCE(rating, 0.0), 
               created_at, updated_at
        FROM movies
        ORDER BY release_date DESC
        LIMIT $1 OFFSET $2
    `

    rows, err := r.db.Query(context.Background(), query, limit, offset)
    if err != nil {
        return nil, 0, fmt.Errorf("error querying movies: %w", err)
    }
    defer rows.Close()
    
    var movies []model.Movie
    for rows.Next() {
        var movie model.Movie
        if err := rows.Scan(
            &movie.MovieID,
            &movie.Title,
            &movie.Description,
            &movie.Duration,
            &movie.ReleaseDate,
            &movie.PosterURL,
            &movie.BackdropURL,
            &movie.Rating,
            &movie.CreatedAt,
            &movie.UpdatedAt,
        ); err != nil {
            return nil, 0, fmt.Errorf("error scanning movie row: %w", err)
        }

        // Fetch genres for this movie
        genres, err := r.GetMovieGenres(movie.MovieID)
        if err != nil {
            return nil, 0, fmt.Errorf("error fetching genres for movie ID %d: %w", movie.MovieID, err)
        }
        movie.Genres = genres

        movies = append(movies, movie)
    }
    
    if err := rows.Err(); err != nil {
        return nil, 0, fmt.Errorf("error iterating movie rows: %w", err)
    }

    // Return empty slice instead of nil if no movies found
    if movies == nil {
        movies = []model.Movie{}
    }

    return movies, totalCount, nil
}

// GetMovieByID retrieves a specific movie by its ID
func (r *MovieRepository) GetMovieByID(id int) (model.Movie, error) {
    query := `
        SELECT movie_id, title, description, duration, release_date, 
               poster_url, backdrop_url, rating, created_at, updated_at
        FROM movies
        WHERE movie_id = $1
    `
    
    var movie model.Movie
    err := r.db.QueryRow(context.Background(), query, id).Scan(
        &movie.MovieID,
        &movie.Title,
        &movie.Description,
        &movie.Duration,
        &movie.ReleaseDate,
        &movie.PosterURL,
        &movie.BackdropURL,
        &movie.Rating,
        &movie.CreatedAt,
        &movie.UpdatedAt,
    )
    
    if err != nil {
        return model.Movie{}, fmt.Errorf("error fetching movie by id: %w", err)
    }
    
    // Fetch genres for this movie
    genres, err := r.GetMovieGenres(movie.MovieID)
    if err != nil {
        return model.Movie{}, err
    }
    movie.Genres = genres
    
    return movie, nil
}

// GetUpcomingMovies retrieves movies with release dates in the future
func (r *MovieRepository) GetUpcomingMovies() ([]model.Movie, error) {
    query := `
        SELECT movie_id, title, COALESCE(description, ''), COALESCE(duration, 0), 
               COALESCE(release_date, CURRENT_DATE), COALESCE(poster_url, ''), 
               COALESCE(backdrop_url, ''), COALESCE(rating, 0.0), 
               created_at, updated_at
        FROM movies
        WHERE release_date > CURRENT_DATE
        ORDER BY release_date ASC
    `
    
    rows, err := r.db.Query(context.Background(), query)
    if err != nil {
        return nil, fmt.Errorf("error querying upcoming movies: %w", err)
    }
    defer rows.Close()
    
    movies := []model.Movie{} // Initialize as empty slice
    for rows.Next() {
        var movie model.Movie
        if err := rows.Scan(
            &movie.MovieID,
            &movie.Title,
            &movie.Description,
            &movie.Duration,
            &movie.ReleaseDate,
            &movie.PosterURL,
            &movie.BackdropURL,
            &movie.Rating,
            &movie.CreatedAt,
            &movie.UpdatedAt,
        ); err != nil {
            return nil, fmt.Errorf("error scanning upcoming movie row: %w", err)
        }
        
        // Fetch genres for this movie
        genres, err := r.GetMovieGenres(movie.MovieID)
        if err != nil {
            return nil, fmt.Errorf("error fetching genres for movie ID %d: %w", movie.MovieID, err)
        }
        movie.Genres = genres
        
        movies = append(movies, movie)
    }
    
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error iterating upcoming movie rows: %w", err)
    }
    
    return movies, nil
}