package users

import "github.com/labstack/echo/v4"

type HttpHandlers interface {
	SignUp() echo.HandlerFunc
}
