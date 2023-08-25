package http

import (
	"avito-test/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type Routes struct {
	Usecase *usecase.UseCase
}

func (routes *Routes) RegisterRoutes(app fiber.Router) {
	app.Post("/user", routes.checkUserSegments)

	app.Post("/segments", routes.createSegments)

	app.Delete("/segments", routes.deleteSegments)
}
