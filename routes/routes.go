package routes

import (
	c "weather-apps/controllers"

	"weather-apps/middlewares"

	"github.com/labstack/echo"
	m "github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	// user routing
	e.GET("/api/users", c.GetUsersController)
	e.GET("/api/users/:id", c.GetUserController)
	e.POST("/api/users", c.CreateUserController)
	e.DELETE("/api/users/:id", c.DeleteUserController, m.BasicAuth(middlewares.BasicAuth2))
	e.PUT("/api/users/:id", c.UpdateUserController, m.BasicAuth(middlewares.BasicAuth2))

	// weather routing
	// e.GET("/api/weather/:id", c.GetWeatherController)
	e.GET("/api/weather/:name", c.GetWeatherController, m.BasicAuth(middlewares.BasicAuth2))

	return e
}
