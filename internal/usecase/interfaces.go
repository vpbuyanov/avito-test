package usecase

import (
	"avito-test/internal/entity"
	"github.com/gofiber/fiber/v2"
)

type (
	Repo interface {
		RepoCreateSegments(ctx *fiber.Ctx, request *entity.ChangeSegmentsRequest, response *entity.ChangeSegmentsResponse) (*entity.ChangeSegmentsResponse, error)
		RepoDeleteSegments(ctx *fiber.Ctx, request *entity.ChangeSegmentsRequest, response *entity.ChangeSegmentsResponse) (*entity.ChangeSegmentsResponse, error)
		RepoCreateUser(ctx *fiber.Ctx, request *entity.CreateUserRequest, response *entity.CreateUserResponse) (*entity.CreateUserResponse, error)
		RepoSetUserSegments(ctx *fiber.Ctx, request *entity.ChangeUserSegmentsRequest, response *entity.ChangeUserSegmentsResponse) (*entity.ChangeUserSegmentsResponse, error)
		RepoGetActiveUserSegments(ctx *fiber.Ctx, request *entity.CheckSegmentsRequest, response *entity.CheckSegmentsResponse) (*entity.CheckSegmentsResponse, error)
	}
	UC interface {
		UseCaseCreateSegments(ctx *fiber.Ctx, request *entity.ChangeSegmentsRequest, response *entity.ChangeSegmentsResponse) (*entity.ChangeSegmentsResponse, error)
		UseCaseDeleteSegments(ctx *fiber.Ctx, request *entity.ChangeSegmentsRequest, response *entity.ChangeSegmentsResponse) (*entity.ChangeSegmentsResponse, error)
		UseCaseGetActiveUserSegments(ctx *fiber.Ctx, request *entity.CheckSegmentsRequest, response *entity.CheckSegmentsResponse) (*entity.CheckSegmentsResponse, error)
		UseCaseCreateUser(ctx *fiber.Ctx, request *entity.CreateUserRequest, response *entity.CreateUserResponse) (*entity.CreateUserResponse, error)
		UseCaseSetUserSegments(ctx *fiber.Ctx, request *entity.ChangeUserSegmentsRequest, response *entity.ChangeUserSegmentsResponse) (*entity.ChangeUserSegmentsResponse, error)
	}
)
