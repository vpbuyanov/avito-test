package usecase

import (
	"avito-test/internal/entity"
	"github.com/gofiber/fiber/v2"
)

type (
	Repo interface {
		RepoCreateSegments(ctx *fiber.Ctx, request *entity.ChangeSegmentsRequest, response *entity.ChangeSegmentsResponse) (*entity.ChangeSegmentsResponse, error)
		RepoDeleteSegments(ctx *fiber.Ctx, request *entity.ChangeSegmentsRequest, response *entity.ChangeSegmentsResponse) (*entity.ChangeSegmentsResponse, error)
		RepoGetActiveUserSegments(ctx *fiber.Ctx, request *entity.CheckSegmentsRequest, response *entity.CheckSegmentsResponse) (*entity.CheckSegmentsResponse, error)
	}
)
