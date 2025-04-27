package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mhakimsaputra17/tickitz-api-go/internal/repository"
)

// AdminHandler handles admin-related endpoints
type AdminHandler struct {
    movieRepo *repository.MovieRepository
}

// NewAdminHandler creates a new AdminHandler instance
func NewAdminHandler(movieRepo *repository.MovieRepository) *AdminHandler {
    return &AdminHandler{
        movieRepo: movieRepo,
    }
}

// GetMovies handles the GET /admin/movies endpoint
func (h *AdminHandler) GetMovies(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "Dummy Admin GetMovies endpoint",
        "status": "success",
    })
}

// CreateMovie handles the POST /admin/movies endpoint
func (h *AdminHandler) CreateMovie(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "Dummy CreateMovie endpoint",
        "status": "success",
    })
}

// UpdateMovie handles the PUT /admin/movies/:id endpoint
func (h *AdminHandler) UpdateMovie(c *gin.Context) {
    id := c.Param("id")
    c.JSON(http.StatusOK, gin.H{
        "message": "Dummy UpdateMovie endpoint",
        "id": id,
        "status": "success",
    })
}

// DeleteMovie handles the DELETE /admin/movies/:id endpoint
func (h *AdminHandler) DeleteMovie(c *gin.Context) {
    id := c.Param("id")
    c.JSON(http.StatusOK, gin.H{
        "message": "Dummy DeleteMovie endpoint",
        "id": id,
        "status": "success",
    })
}