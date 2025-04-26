package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mhakimsaputra17/tickitz-api-go/internal/model"
	"github.com/mhakimsaputra17/tickitz-api-go/internal/repository"
)

// AuthHandler handles authentication requests
type AuthHandler struct {
	userRepo *repository.UserRepository
}

// NewAuthHandler creates a new AuthHandler instance
func NewAuthHandler (userRepo *repository.UserRepository) *AuthHandler {
	return &AuthHandler{
		userRepo: userRepo,
	}
}

// Register handles user registration

func (h *AuthHandler) Register(c *gin.Context){
	var input model.UserRegister

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	// Check if user already exists
	_, err := h.userRepo.GetUserByEmail(c, input.Email)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		return
	}

	user, err := h.userRepo.CreateUser(c, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, model.AuthResponse{
		User: user,
		Message : "User registered successfully",
	})
}

// Login handles user authentication
func (h *AuthHandler) Login (c *gin.Context){
	var input model.UserLogin

	if err:= c.ShouldBindJSON(&input); err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return 
	}

	// Get user from database
	user, err := h.userRepo.GetUserByEmail(c, input.Email)
	
	if err!= nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error" : "Invalid credentials",
		})
		return
	}

	// Verify password 
	if !h.userRepo.VerifyPassword(user.PasswordHash, input.PasswordHash){
		c.JSON(http.StatusUnauthorized, gin.H{
			"error" :"Invalid credentials",
		})
		return
	}

	// Clear password field for response
	c.JSON(http.StatusOK, model.AuthResponse{
		User: user,
		Message: "Login successful",
	})


}


