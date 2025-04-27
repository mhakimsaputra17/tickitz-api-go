package handler

import (
	// "fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mhakimsaputra17/tickitz-api-go/internal/repository"
	"github.com/mhakimsaputra17/tickitz-api-go/internal/util"
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
     // Parse and validate pagination parameters
	 page, err:= strconv.Atoi(c.DefaultQuery("page", "1"))
	 if err != nil || page < 1 {
		page = 1
	 }

	 limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	 if err != nil || limit < 1 || limit > 50 {
		limit = 10
	 }

	 // Call repository to get movies with genres
	 movies, totalCount, err := h.movieRepo.GetMovies(page, limit)
	 if err != nil {
		util.ServerErrorResponse(c, "Failed to fetch movies")
		return
	 }
	//  fmt.Println(movies)

	 // Create pagination metadata
	 totalPages := (totalCount + limit - 1) / limit // Ceiling division

	 pagination := map[string] interface{}{
		"current_page": page,
        "total_pages":  totalPages,
        "limit":        limit,
        "total_items":  totalCount,
}
	response := map[string] interface{}{
		"movies" : movies,
		"pagination" : pagination,
	}

	util.OkResponse(c, "Movies retrieved successfully", response)

}



// GetUpcomingMovies handles the GET /movies/upcoming endpoint
func (h *MovieHandler) GetUpcomingMovies(c *gin.Context) {
    movies, err := h.movieRepo.GetUpcomingMovies()
	// fmt.Println(movies)
    if err != nil {
        util.ServerErrorResponse(c, "Failed to fetch upcoming movies")
        return
    }
    
    util.OkResponse(c, "Upcoming movies retrieved successfully", movies)
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
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        util.BadRequestResponse(c, "Invalid movie ID format", nil)
        return
    }
    
    movie, err := h.movieRepo.GetMovieByID(id)
    if err != nil {
        // Check for no rows / not found error
        if err.Error() == "error fetching movie by id: no rows in result set" {
            util.NotFoundResponse(c, "Movie not found")
            return
        }
        util.ServerErrorResponse(c, "Failed to fetch movie details")
        return
    }
    
    util.OkResponse(c, "Movie details retrieved successfully", movie)
}