package routes

import (
	"github.com/clim-bot/url-shortener/controllers"
	"github.com/clim-bot/url-shortener/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())

	AuthRoutes(r)
	URLRoutes(r)
	r.GET("/health", controllers.HealthCheck) // Add this line

	return r
}
