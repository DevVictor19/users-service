package http

import (
	"usersservice/internal/users"

	"github.com/labstack/echo/v4"
)

func MapUsersRoutes(usersGroup *echo.Group, usersHandlers users.HttpHandlers) {
	usersGroup.POST("/auth/signup", usersHandlers.SignUp())
}
