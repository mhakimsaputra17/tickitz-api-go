package handler

import "github.com/gin-gonic/gin"

// AuthHandlerInterface defines methods for authentication handlers

type AuthHandlerInterface interface {
	Register (c *gin.Context)
	Login(c *gin.Context)
}


// MovieHandlerInterface defines methods for movie handlers

type MovieHandlerInterface interface {
	GetMovies (c *gin.Context)
	GetUpcomingMovies(c *gin.Context)
	GetPopularMovies(c *gin.Context)
	GetMovieByID(c *gin.Context)
}

// ScheduleHandlerInterface defines methods for schedule handlers
type UserHandlerInterface interface {
	GetProfile (c *gin.Context)
	UpdateProfile (c *gin.Context)
	GetOrders (c *gin.Context)
	CreateOrder (c *gin.Context)
}

// AdminHandlerInterface defines methods for admin handlers

type AdminHandlerInterface interface {
	GetMovies(c *gin.Context)
	CreateMovie(c *gin.Context)
	UpdateMovie(c *gin.Context)
	DeleteMovie(c *gin.Context)
}