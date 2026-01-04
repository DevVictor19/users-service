package server

import (
	"net/http"
	_ "usersservice/docs"
	usersHttpHandlers "usersservice/internal/users/interfaces/http"
	usersRepositories "usersservice/internal/users/repositories"
	usersUseCase "usersservice/internal/users/usecase"
	"usersservice/pkg/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func (s *Server) MapHandlers(e *echo.Echo) error {
	// Init repositories
	userRepo := usersRepositories.NewMySQLUsersRepository(s.db)

	// Init useCases
	userUC := usersUseCase.NewUsersUseCase(s.cfg, userRepo, s.logger)

	// Init handlers
	userHandlers := usersHttpHandlers.NewUsersHandlers(s.cfg, userUC, s.logger)

	// Middlewares
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         1 << 10, // 1 KB
		DisablePrintStack: true,
		DisableStackAll:   true,
	}))
	e.Use(middleware.RequestID())
	e.Use(middleware.Secure())
	e.Use(middleware.BodyLimit("2M"))
	e.Use(makeRequestLoggerMiddleware(s.logger))

	// Routes
	v1 := e.Group("/api/v1")

	v1.GET("/health", func(c echo.Context) error {
		s.logger.Infof("Health check RequestID: %s", utils.GetRequestID(c))
		return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
	})

	v1.GET("/swagger/*", echoSwagger.WrapHandler)

	// Map handlers
	usersHttpHandlers.MapUsersRoutes(v1, userHandlers)

	return nil
}
