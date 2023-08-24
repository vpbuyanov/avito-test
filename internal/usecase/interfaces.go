package usecase

import (
	"avito-test/internal/entity"
	"github.com/gofiber/fiber/v2"
)

type (
	Repo interface {
		RepoCreateSegments(ctx *fiber.Ctx, request *entity.ChangeSegmentsRequest, response *entity.ChangeSegmentsResponse) (*entity.ChangeSegmentsResponse, error)
	}
)
