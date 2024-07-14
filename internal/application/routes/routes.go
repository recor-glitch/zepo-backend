package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/recor-glitch/zepo-backend/internal/application/handlers"
)

func SetupRoutes(r *gin.Engine) {
	// USER
	r.GET("/user/:id", handlers.GetUserByID)
	r.POST("/user", handlers.CreateUser)
}