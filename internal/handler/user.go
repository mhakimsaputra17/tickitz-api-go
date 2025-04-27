package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mhakimsaputra17/tickitz-api-go/internal/repository"
)

// UserHandler handles user-related endpoints
type UserHandler struct {
    userRepo *repository.UserRepository
}

// NewUserHandler creates a new UserHandler instance
func NewUserHandler(userRepo *repository.UserRepository) *UserHandler {
    return &UserHandler{
        userRepo: userRepo,
    }
}

// GetProfile handles the GET /user/profile endpoint
func (h *UserHandler) GetProfile(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "Dummy GetProfile endpoint",
        "status": "success",
    })
}

// UpdateProfile handles the PUT /user/profile endpoint
func (h *UserHandler) UpdateProfile(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "Dummy UpdateProfile endpoint",
        "status": "success",
    })
}

// GetOrders handles the GET /user/orders endpoint
func (h *UserHandler) GetOrders(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "Dummy GetOrders endpoint",
        "status": "success",
    })
}

// CreateOrder handles the POST /user/orders endpoint
func (h *UserHandler) CreateOrder(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "Dummy CreateOrder endpoint",
        "status": "success",
    })
}


