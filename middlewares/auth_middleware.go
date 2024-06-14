package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/clim-bot/url-shortener/utils"
	"github.com/gin-gonic/gin"
)

var ctx = context.Background()

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Check if token is blacklisted
		if utils.IsTokenBlacklisted(ctx, tokenString) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token has been invalidated"})
			c.Abort()
			return
		}

		claims, valid := utils.ValidateToken(tokenString)
		if !valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Next()
	}
}
