package usecase

import (
	"avito-test/internal/entity"
	"github.com/gofiber/fiber/v2"
)

type UseCase struct {
	repo Repo
}

func NewUseCase(r Repo) *UseCase {
	return &UseCase{
		repo: r,
	}
}

func (use *UseCase) UseCaseCreateSegments(ctx *fiber.Ctx, request *entity.ChangeSegmentsRequest, response *entity.ChangeSegmentsResponse) (*entity.ChangeSegmentsResponse, error) {
	response, err := use.repo.RepoCreateSegments(ctx, request, response)
	if err != nil {
		response.Status = "Error"
		return response, nil
	}

	return response, nil
}

func (use *UseCase) UseCaseDeleteSegments(ctx *fiber.Ctx, request *entity.ChangeSegmentsRequest, response *entity.ChangeSegmentsResponse) (*entity.ChangeSegmentsResponse, error) {
	response, err := use.repo.RepoDeleteSegments(ctx, request, response)
	if err != nil {
		response.Status = "Error"
		return response, nil
	}

	return response, nil
}

func (use *UseCase) UseCaseGetActiveUserSegments(ctx *fiber.Ctx, request *entity.CheckSegmentsRequest, response *entity.CheckSegmentsResponse) (*entity.CheckSegmentsResponse, error) {
	response, err := use.repo.RepoGetActiveUserSegments(ctx, request, response)
	if err != nil {
		response.Status = "Error"
		return response, nil
	}

	return response, nil
}

func (use *UseCase) UseCaseCreateUser(ctx *fiber.Ctx, request *entity.CreateUserRequest, response *entity.CreateUserResponse) (*entity.CreateUserResponse, error) {
	response, err := use.repo.RepoCreateUser(ctx, request, response)

	if err != nil {
		response.Status = "Error"
		return response, nil
	}

	return response, nil
}

func (use *UseCase) UseCaseSetUserSegments(ctx *fiber.Ctx, request *entity.ChangeUserSegmentsRequest, response *entity.ChangeUserSegmentsResponse) (*entity.ChangeUserSegmentsResponse, error) {
	response, err := use.repo.RepoSetUserSegments(ctx, request, response)

	if err != nil {
		response.Status = "Error"
		return response, nil
	}

	return response, nil
}
