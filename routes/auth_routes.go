package routes

import (
	"github.com/clim-bot/url-shortener/controllers"
	"github.com/clim-bot/url-shortener/middlewares"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
		auth.POST("/logout", middlewares.AuthMiddleware(), controllers.Logout)
	}
}
