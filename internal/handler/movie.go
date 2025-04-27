package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mhakimsaputra17/tickitz-api-go/internal/repository"
)

// MovieHandler handles movie-related endpoints
type MovieHandler struct {
    movieRepo *repository.MovieRepository
}

// NewMovieHandler creates a new MovieHandler instance
func NewMovieHandler(movieRepo *repository.MovieRepository) *MovieHandler {
    return &MovieHandler{
        movieRepo: movieRepo,
    }
}

// GetMovies handles the GET /movies endpoint
func (h *MovieHandler) GetMovies(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "Dummy GetMovies endpoint",
        "status": "success",
    })
}

// GetUpcomingMovies handles the GET /movies/upcoming endpoint
func (h *MovieHandler) GetUpcomingMovies(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "Dummy GetUpcomingMovies endpoint",
        "status": "success",
    })
}

// GetPopularMovies handles the GET /movies/popular endpoint
func (h *MovieHandler) GetPopularMovies(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "Dummy GetPopularMovies endpoint",
        "status": "success",
    })
}

// GetMovieByID handles the GET /movies/:id endpoint
func (h *MovieHandler) GetMovieByID(c *gin.Context) {
    id := c.Param("id")
    c.JSON(http.StatusOK, gin.H{
        "message": "Dummy GetMovieByID endpoint",
        "id": id,
        "status": "success",
    })
}