package http

import (
	"avito-test/internal/entity"
	"github.com/gofiber/fiber/v2"
)

func (routes *Routes) createSegments(ctx *fiber.Ctx) error {
	request := new(entity.ChangeSegmentsRequest)
	response := new(entity.ChangeSegmentsResponse)

	err := ctx.BodyParser(request)
	if err != nil {
		return err
	}

	response, err = routes.Usecase.UseCaseCreateSegments(ctx, request, response)
	if err != nil {
		return err
	}

	return ctx.JSON(response)
}

func (routes *Routes) deleteSegments(ctx *fiber.Ctx) error {
	request := new(entity.ChangeSegmentsRequest)
	response := new(entity.ChangeSegmentsResponse)

	err := ctx.BodyParser(request)
	if err != nil {
		return err
	}

	response, err = routes.Usecase.UseCaseDeleteSegments(ctx, request, response)
	if err != nil {
		return err
	}

	return ctx.JSON(response)
}
