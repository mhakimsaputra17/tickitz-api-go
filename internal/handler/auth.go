package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mhakimsaputra17/tickitz-api-go/internal/model"
	"github.com/mhakimsaputra17/tickitz-api-go/internal/repository"
)

// AuthHandler handles authentication-related endpoints
type AuthHandler struct {
	userRepo *repository.UserRepository
}

// NewAuthHandler creates a new AuthHandler instance
func NewAuthHandler(userRepo *repository.UserRepository) *AuthHandler {
	return &AuthHandler{
		userRepo: userRepo,
	}
}

// Register handles user registration
func (h *AuthHandler) Register(c *gin.Context) {
	var input model.UserRegister

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"message": err.Error(),
		})
		return
	}

	// Check if user already exists
	_, err := h.userRepo.GetUserByEmail(c, input.Email)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"status": "error",
			"message": "Email already exists",
		})
		return
	}

	user, err := h.userRepo.CreateUser(c, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"message": "Failed to create user",
		})
		return
	}

	// Build response according to API spec
	response := model.AuthResponse{
		Status: "success",
		Data: struct {
			User      model.User `json:"user"`
			Token     string     `json:"token,omitempty"`
			TokenType string     `json:"token_type,omitempty"`
			ExpiresIn int        `json:"expires_in,omitempty"`
		}{
			User: user,
		},
	}

	c.JSON(http.StatusCreated, response)
}

// Login handles user authentication
func (h *AuthHandler) Login(c *gin.Context) {
	var input model.UserLogin

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"message": err.Error(),
		})
		return
	}

	// Get user from database
	user, err := h.userRepo.GetUserByEmail(c, input.Email)
	
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "error",
			"message": "Invalid credentials",
		})
		return
	}

	// Verify password 
	if !h.userRepo.VerifyPassword(user.PasswordHash, input.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "error",
			"message": "Invalid credentials",
		})
		return
	}

	// TODO: In the future, you'd generate a JWT token here

	// Build response according to API spec
	response := model.AuthResponse{
		Status: "success",
		Data: struct {
			User      model.User `json:"user"`
			Token     string     `json:"token,omitempty"`
			TokenType string     `json:"token_type,omitempty"`
			ExpiresIn int        `json:"expires_in,omitempty"`
		}{
			User:      user,
			Token:     "dummy-token", // Replace with actual JWT token
			TokenType: "Bearer",
			ExpiresIn: 86400, // 24 hours in seconds
		},
	}

	c.JSON(http.StatusOK, response)
}
