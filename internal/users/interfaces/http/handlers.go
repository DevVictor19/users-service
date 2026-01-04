package http

import (
	"net/http"
	"usersservice/internal/users"
	"usersservice/internal/users/dtos"
	"usersservice/pkg/env"
	"usersservice/pkg/httpErrors"
	"usersservice/pkg/logger"
	"usersservice/pkg/utils"

	"github.com/labstack/echo/v4"
)

type usersHandlers struct {
	cfg     *env.Config
	usersUC users.UseCase
	logger  logger.Logger
}

func NewUsersHandlers(cfg *env.Config, usersUC users.UseCase, logger logger.Logger) users.HttpHandlers {
	return &usersHandlers{
		cfg:     cfg,
		usersUC: usersUC,
		logger:  logger,
	}
}

// SignUp godoc
// @Summary      Sign Up
// @Description  Register a new user
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        input  body      dtos.SignUpInputDTO  true  "Sign Up Input"
// @Success      201    {object}  dtos.UserDTO
// @Failure      400    {object}  httpErrors.RestError
// @Failure      500    {object}  httpErrors.RestError
// @Router       /users/auth/signup [post]
func (h *usersHandlers) SignUp() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := &dtos.SignUpInputDTO{}
		if err := c.Bind(input); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(http.StatusBadRequest, httpErrors.NewRestErrorWithMessage(
				http.StatusBadRequest,
				"Invalid body",
				err,
			))
		}

		output, err := h.usersUC.SignUp(utils.GetRequestCtx(c), input)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.NewFromError(err))
		}

		return c.JSON(http.StatusCreated, output)
	}
}
