package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mhakimsaputra17/tickitz-api-go/internal/model"
	"github.com/mhakimsaputra17/tickitz-api-go/internal/repository"
	"github.com/mhakimsaputra17/tickitz-api-go/internal/util"
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
		util.BadRequestResponse(c, "Invalid input", err.Error())
		return
	}

	// Check if user already exists
	_, err := h.userRepo.GetUserByEmail(c, input.Email)
	if err == nil {
		util.ConflictResponse(c, "Email already exists")
		return
	}

	user, err := h.userRepo.CreateUser(c, input)
	if err != nil {
		util.ServerErrorResponse(c, "Failed to create user")
		return
	}

	// Struktur data respons
	responseData := struct {
		User model.User `json:"user"`
	}{
		User: user,
	}

	util.CreatedResponse(c, "User registered successfully", responseData)
}

// Login handles user authentication
func (h *AuthHandler) Login(c *gin.Context) {
	var input model.UserLogin

	if err := c.ShouldBindJSON(&input); err != nil {
		util.BadRequestResponse(c, "Invalid input", err.Error())
		return
	}

	// Get user from database
	user, err := h.userRepo.GetUserByEmail(c, input.Email)
	
	if err != nil {
		util.UnauthorizedResponse(c, "Invalid credentials")
		return
	}

	// Verify password 
	if !h.userRepo.VerifyPassword(user.PasswordHash, input.Password) {
		util.UnauthorizedResponse(c, "Invalid credentials")
		return
	}

	// Struktur data respons
	responseData := struct {
		User      model.User `json:"user"`
		Token     string     `json:"token"`
		TokenType string     `json:"token_type"`
		ExpiresIn int        `json:"expires_in"`
	}{
		User:      user,
		Token:     "dummy-token", // Ganti dengan JWT asli
		TokenType: "Bearer",
		ExpiresIn: 86400, // 24 jam dalam detik
	}

	util.OkResponse(c, "Login successful", responseData)
}
