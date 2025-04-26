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

func main(){
 
	// Load .env
	if err := godotenv.Load("./config/.env"); err!= nil {
		log.Println("No .env file found")
	}

	// Load config
	cfg:= config.LoadConfig()

	// Init DB
	dbPool, err:= database.NewPostgresPool(cfg)
	if err != nil {
		 log.Fatalf("Failed to connect to DB: %v", err)
	}

	defer dbPool.Close()

	// Init repository & handler (dependency injection)
	userRepo := repository.NewUserRepository(dbPool)
	authHandler := handler.NewAuthHandler(userRepo)




	r := gin.Default()
	router.SetupRoutes(r, authHandler)
	r.Run()



}