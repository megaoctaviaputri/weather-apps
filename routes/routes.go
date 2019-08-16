package routes

import (
	c "go-restful-training/controllers"
	"go-restful-training/middlewares"

	"github.com/labstack/echo"
	m "github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	// user routing
	e.GET("/api/users", c.GetUsersController, m.BasicAuth(middlewares.BasicAuth2))
	e.GET("/api/users/:id", c.GetUserController, m.BasicAuth(middlewares.BasicAuth))
	e.POST("/api/users", c.CreateUserController)
	e.DELETE("/api/users/:id", c.DeleteUserController, m.BasicAuth(middlewares.BasicAuth2))
	e.PUT("/api/users/:id", c.UpdateUserController, m.BasicAuth(middlewares.BasicAuth2))

	return e
}
