package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mhakimsaputra17/tickitz-api-go/internal/handler"
)

func SetupRoutes(r *gin.Engine, authHandler *handler.AuthHandler){
	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)
}