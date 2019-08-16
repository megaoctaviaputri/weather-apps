package routes

import (
	c "weather-apps/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	// user routing
	e.GET("/api/users", c.GetUsersController)
	e.GET("/api/users/:id", c.GetUserController)
	e.POST("/api/users", c.CreateUserController)
	e.DELETE("/api/users/:id", c.DeleteUserController)
	e.PUT("/api/users/:id", c.UpdateUserController)

	// weather routing
	// e.GET("/api/weather/:id", c.GetWeatherController)
	e.GET("/api/weather/:name", c.GetWeatherController)

	return e
}
