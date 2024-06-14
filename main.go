package main

import (
	"github.com/clim-bot/url-shortener/config"
	"github.com/clim-bot/url-shortener/models"
	"github.com/clim-bot/url-shortener/routes"
)

func main() {
	config.Init()

	// Drop the table if it exists and recreate it (useful during development)
	// config.DB.Migrator().DropTable(&models.URL{})
	// Auto migrate the models
	config.DB.AutoMigrate(&models.User{}, &models.URL{})

	r := routes.SetupRouter()
	r.Run()
}
