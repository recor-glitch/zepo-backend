package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/recor-glitch/zepo-backend/internal/application/handlers"
	"github.com/recor-glitch/zepo-backend/internal/application/middleware"
)

func SetupRoutes(r *gin.Engine) {
	// USER ROUTES
	// INVALIDATE TOKEN
	r.POST("/invalidate", handlers.InvalidateAccessToken)
	// CREATE USER
	r.POST("/user", handlers.CreateUser)
	r.POST("/get-by-email", handlers.GetUserByEmail)
	authRoute := r.Group("/user")

	authRoute.Use(middleware.JWTAuth())

	authRoute.GET("/:id", handlers.GetUserByID)
}
