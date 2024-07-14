package server

import (
	"github.com/gin-gonic/gin"
	"github.com/recor-glitch/zepo-backend/internal/application/routes"
)

func Run() {
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}