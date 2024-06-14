package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/clim-bot/url-shortener/config"
	"github.com/clim-bot/url-shortener/models"
	"github.com/clim-bot/url-shortener/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateShortURL(c *gin.Context) {
	var request struct {
		OriginalURL string `json:"original_url" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username, _ := c.Get("username")
	var user models.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}

	shortCode := uuid.New().String()[:6]

	// Check if the short code already exists
	var existingURL models.URL
	if err := config.DB.Where("short_code = ?", shortCode).First(&existingURL).Error; err == nil {
		log.Printf("Short code already exists: %s", shortCode)
		c.JSON(http.StatusConflict, gin.H{"error": "Short code already exists"})
		return
	}

	url := models.URL{
		OriginalURL: request.OriginalURL,
		ShortCode:   shortCode,
		UserID:      user.ID,
	}

	if err := config.DB.Create(&url).Error; err != nil {
		log.Printf("Failed to create short URL: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create short URL"})
		return
	}

	// Set cache
	utils.SetCache(shortCode, url.OriginalURL, 10*time.Minute)

	c.JSON(http.StatusOK, gin.H{"short_code": shortCode})
}

func GetOriginalURL(c *gin.Context) {
	shortCode := c.Param("shortCode")
	
	// Check cache first
	if cachedURL, err := utils.GetCache(shortCode); err == nil {
		c.Redirect(http.StatusMovedPermanently, cachedURL)
		return
	}

	var url models.URL
	if err := config.DB.Where("short_code = ?", shortCode).First(&url).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	// Set cache
	utils.SetCache(shortCode, url.OriginalURL, 10*time.Minute)

	c.Redirect(http.StatusMovedPermanently, url.OriginalURL)
}