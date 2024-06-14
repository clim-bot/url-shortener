package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/clim-bot/url-shortener/config"
	"github.com/clim-bot/url-shortener/models"
	"github.com/clim-bot/url-shortener/utils"
	"github.com/clim-bot/url-shortener/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	config.Init()
	code := m.Run()
	os.Exit(code)
}

func TestRegister(t *testing.T) {
	r := gin.Default()
	r.POST("/auth/register", Register)

	tests := []struct {
		name         string
		user         models.User
		expectedCode int
	}{
		{
			name: "Valid Input",
			user: models.User{Username: "testuser", Password: "password123"},
			expectedCode: http.StatusOK,
		},
		{
			name: "Short Username",
			user: models.User{Username: "tu", Password: "password123"},
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Short Password",
			user: models.User{Username: "testuser", Password: "pass"},
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Invalid Username Characters",
			user: models.User{Username: "test_user!", Password: "password123"},
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonValue, _ := json.Marshal(tt.user)
			req, _ := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(jsonValue))
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
		})
	}
}

func TestLogin(t *testing.T) {
	r := gin.Default()
	r.POST("/auth/login", Login)

	// Create a user for login tests
	user := models.User{Username: "testuser", Password: "password123"}
	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword
	config.DB.Create(&user)

	tests := []struct {
		name         string
		user         models.User
		expectedCode int
	}{
		{
			name: "Valid Input",
			user: models.User{Username: "testuser", Password: "password123"},
			expectedCode: http.StatusOK,
		},
		{
			name: "Short Username",
			user: models.User{Username: "tu", Password: "password123"},
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Short Password",
			user: models.User{Username: "testuser", Password: "pass"},
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Invalid Username Characters",
			user: models.User{Username: "test_user!", Password: "password123"},
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Invalid Credentials",
			user: models.User{Username: "testuser", Password: "wrongpassword"},
			expectedCode: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonValue, _ := json.Marshal(tt.user)
			req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonValue))
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
		})
	}
}

func TestLogout(t *testing.T) {
	r := gin.Default()
	r.POST("/auth/logout", middlewares.AuthMiddleware(), Logout)

	// Create a user for logout tests
	user := models.User{Username: "testuser", Password: "password123"}
	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword
	config.DB.Create(&user)

	token, _ := utils.GenerateJWT(user.Username)

	tests := []struct {
		name         string
		token        string
		expectedCode int
	}{
		{
			name: "Valid Token",
			token: token,
			expectedCode: http.StatusOK,
		},
		{
			name: "Invalid Token",
			token: "invalidtoken",
			expectedCode: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("POST", "/auth/logout", nil)
			req.Header.Set("Authorization", "Bearer "+tt.token)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
		})
	}
}
