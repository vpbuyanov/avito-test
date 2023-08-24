package http

import (
	"avito-test/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type Routes struct {
	Usecase *usecase.UseCase
}

func (routes *Routes) RegisterRoutes(app fiber.Router) {
	app.Post("/segments/create", routes.createSegments)
}
