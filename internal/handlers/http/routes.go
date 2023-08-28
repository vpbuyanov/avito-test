package http

import (
	"avito-test/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type Routes struct {
	Usecase usecase.UC
}

func (routes *Routes) RegisterRoutes(app fiber.Router) {
	app.Post("/user", routes.createUser)
	app.Post("/user/check", routes.checkUserSegments)
	app.Post("/user/change", routes.changeUserSegments)

	app.Post("/segments", routes.createSegments)

	app.Delete("/segments", routes.deleteSegments)
}
