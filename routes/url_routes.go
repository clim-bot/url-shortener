package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/clim-bot/url-shortener/controllers"
	"github.com/clim-bot/url-shortener/middlewares"
)

func URLRoutes(r *gin.Engine) {
	url := r.Group("/url")
	{
		url.POST("/shorten", middlewares.AuthMiddleware(), controllers.CreateShortURL)
		url.GET("/:shortCode", middlewares.RateLimiter(), controllers.GetOriginalURL)
	}
}
