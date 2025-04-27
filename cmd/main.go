package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mhakimsaputra17/tickitz-api-go/config"
	"github.com/mhakimsaputra17/tickitz-api-go/internal/handler"
	"github.com/mhakimsaputra17/tickitz-api-go/internal/repository"
	"github.com/mhakimsaputra17/tickitz-api-go/internal/router"
	"github.com/mhakimsaputra17/tickitz-api-go/pkg/database"
)

func main() {
	// Load .env
	if err := godotenv.Load("../config/.env"); err != nil {
		log.Println("No .env file found")
	}

	// Load config
	cfg := config.LoadConfig()

	// Init DB
	dbPool, err := database.NewPostgresPool(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	defer dbPool.Close()

	// Init repositories
	userRepo := repository.NewUserRepository(dbPool)
	movieRepo := repository.NewMovieRepository(dbPool)
	scheduleRepo := repository.NewScheduleRepository(dbPool)
	

	// Init handlers
	authHandler := handler.NewAuthHandler(userRepo)
	movieHandler := handler.NewMovieHandler(movieRepo)
	scheduleHandler := handler.NewScheduleHandler(scheduleRepo)
	adminHandler := handler.NewAdminHandler(movieRepo)
	userHandler := handler.NewUserHandler(userRepo)

	// Setup router
	r := gin.Default()
	router.SetupRoutes(r, authHandler, movieHandler, scheduleHandler, adminHandler, userHandler)
	
	// Run server
	r.Run(":8080")
}