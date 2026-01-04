package http

import (
	"usersservice/internal/users"

	"github.com/labstack/echo/v4"
)

func MapUsersRoutes(router *echo.Group, usersHandlers users.HttpHandlers) {
	usersGroup := router.Group("/users")
	usersGroup.POST("/auth/signup", usersHandlers.SignUp())
}
