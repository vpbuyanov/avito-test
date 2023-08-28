package server

import (
	"avito-test/config"
	"avito-test/internal/handlers/http"
	"avito-test/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	config  *config.Config
	usecase usecase.UC
}

func NewServer(config config.Config, useCase usecase.UC) *Server {
	return &Server{
		config:  &config,
		usecase: useCase,
	}
}

func (s *Server) Run() error {
	app := fiber.New()
	app.Use(logger.New())

	api := app.Group("/api")

	routes := http.Routes{Usecase: s.usecase}
	routes.RegisterRoutes(api)

	return app.Listen(s.config.Server.Address)
}
