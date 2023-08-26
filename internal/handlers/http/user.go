package http

import (
	"avito-test/internal/entity"
	"github.com/gofiber/fiber/v2"
)

func (routes *Routes) createUser(ctx *fiber.Ctx) error {
	request := new(entity.CreateUserRequest)
	response := new(entity.CreateUserResponse)

	err := ctx.BodyParser(request)
	if err != nil {
		return err
	}

	response, err = routes.Usecase.UseCaseCreateUser(ctx, request, response)
	if err != nil {
		return err
	}

	return ctx.JSON(response)
}

func (routes *Routes) checkUserSegments(ctx *fiber.Ctx) error {
	request := new(entity.CheckSegmentsRequest)
	response := new(entity.CheckSegmentsResponse)

	err := ctx.BodyParser(request)
	if err != nil {
		return err
	}

	response, err = routes.Usecase.UseCaseGetActiveUserSegments(ctx, request, response)
	if err != nil {
		return err
	}

	return ctx.JSON(response)
}

func (routes *Routes) changeUserSegments(ctx *fiber.Ctx) error {
	request := new(entity.ChangeUserSegmentsRequest)
	response := new(entity.ChangeUserSegmentsResponse)

	err := ctx.BodyParser(request)
	if err != nil {
		return err
	}

	response, err = routes.Usecase.UseCaseSetUserSegments(ctx, request, response)
	if err != nil {
		return err
	}

	return ctx.JSON(response)
}
