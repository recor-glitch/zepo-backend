package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/recor-glitch/zepo-backend/internal/application/handlers"
	"github.com/recor-glitch/zepo-backend/internal/application/middleware"
)

func SetupRoutes(r *gin.Engine) {
	// USER ROUTES
	// CREATE USER
	r.POST("/user", handlers.CreateUser)
	authRoute := r.Group("/user")

	authRoute.Use(middleware.JWTAuth())

	authRoute.GET("/:id", handlers.GetUserByID)
}
