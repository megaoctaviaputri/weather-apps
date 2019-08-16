package middlewares

import (
	"weather-apps/config"
	"weather-apps/controllers"
	m "weather-apps/models"

	"github.com/labstack/echo"
)

func BasicAuth2(username, password string, c echo.Context) (bool, error) {
	var db = config.DB
	var user m.User
	password = controllers.ToMD5(password)
	if err := db.Where("email = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return false, nil
	}
	return true, nil
}
