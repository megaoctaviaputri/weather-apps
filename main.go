package main

import (
	"go-restful-training/config"
	"go-restful-training/models"
	"go-restful-training/routes"

	m "go-restful-training/middlewares"
)

func main() {
	e := routes.New()

	// implement middleware
	m.LogMiddlewares(e)
	// migration database
	// InitialMigration()

	e.Start(":8000")
}

func InitialMigration() {
	var db = config.DB
	if !db.HasTable(&models.User{}) { // check database exist or not
		db.AutoMigrate(&models.User{})
	}
}
