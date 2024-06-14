package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/clim-bot/url-shortener/config"
	"github.com/clim-bot/url-shortener/models"
	"github.com/clim-bot/url-shortener/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateShortURL(t *testing.T) {
	config.Init()
	r := gin.Default()
	r.POST("/url/shorten", CreateShortURL)

	// Create a user for testing
	user := models.User{Username: "testuser", Password: "password123"}
	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword
	config.DB.Create(&user)

	// Simulate logged-in user by setting the context
	tests := []struct {
		name         string
		originalURL  string
		expectedCode int
	}{
		{
			name: "Valid URL",
			originalURL: "http://example.com",
			expectedCode: http.StatusOK,
		},
		{
			name: "Invalid URL",
			originalURL: "",
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonValue, _ := json.Marshal(map[string]string{"original_url": tt.originalURL})
			req, _ := http.NewRequest("POST", "/url/shorten", bytes.NewBuffer(jsonValue))
			req.Header.Set("Authorization", "Bearer valid_token") // Replace with a valid token or mock the middleware
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
		})
	}
}
