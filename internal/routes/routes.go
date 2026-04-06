package routes

import (
	"cv-api/internal/handlers"
	"cv-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.ErrorMiddleware())

	// Basics
	r.GET("/basics", handlers.GetBasics)
	r.PUT("/basics", handlers.UpdateBasics)

	// Work
	r.GET("/experiences", handlers.GetExperiences)
	r.GET("/experiences/:id", handlers.GetExperience)
	r.PUT("/experiences/:id", handlers.UpdateExperience)

	return r
}
