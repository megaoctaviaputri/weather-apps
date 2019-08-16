package main

import (
	"weather-apps/config"
	"weather-apps/models"
	"weather-apps/routes"

	m "weather-apps/middlewares"
)

func main() {
	e := routes.New()

	// implement middleware
	m.LogMiddlewares(e)
	// migration database
	// InitialMigration()

	e.Start(":7000")
}

func InitialMigration() {
	var db = config.DB
	if !db.HasTable(&models.User{}) { // check database exist or not
		db.AutoMigrate(&models.User{})
	}
}
