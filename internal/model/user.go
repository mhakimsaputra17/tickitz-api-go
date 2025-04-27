package model

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID            int       `json:"user_id"`
	Username      string    `json:"username"`
	Email         string    `json:"email"`
	PasswordHash  string    `json:"-"` // Password is never returned in JSON
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	PhoneNumber   string    `json:"phone_number"`
	ProfileImage  string    `json:"profile_image"`
	LoyaltyPoints int       `json:"loyalty_points"`
	Tier          string    `json:"tier"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// UserRegister is the input model for user registration
type UserRegister struct {
	Username            string `json:"username" binding:"required,min=3,max=30"`
	Email               string `json:"email" binding:"required,email"`
	Password            string `json:"password" binding:"required,min=8,max=30"`
	// PasswordConfirmation string `json:"password_confirmation" binding:"required,eqfield=Password"`
	FirstName           string `json:"first_name" `
	LastName            string `json:"last_name" `
	PhoneNumber         string `json:"phone_number"`
}

// UserLogin is the input model for user login
type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// AuthResponse contains the user data and message for authentication responses
type AuthResponse struct {
	Status string `json:"status"`
	Data   struct {
		User      User   `json:"user"`
		Token     string `json:"token,omitempty"`
		TokenType string `json:"token_type,omitempty"`
		ExpiresIn int    `json:"expires_in,omitempty"`
	} `json:"data"`
}