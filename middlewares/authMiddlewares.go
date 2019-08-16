package middlewares

import (
	"go-restful-training/config"
	"go-restful-training/controllers"
	m "go-restful-training/models"

	"github.com/labstack/echo"
)

func BasicAuth(username string, password string, c echo.Context) (bool, error) {
	if username == "admin" && password == "admin" {
		return true, nil
	}
	return false, nil
}

func BasicAuth2(username, password string, c echo.Context) (bool, error) {
	var db = config.DB
	var user m.User
	password = controllers.ToMD5(password)
	if err := db.Where("email = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return false, nil
	}
	return true, nil
}
