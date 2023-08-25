package repos

import (
	"avito-test/internal/databases"
	"avito-test/internal/entity"
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type Repositories struct {
	db *sql.DB
}

func New(db *sql.DB) *Repositories {
	return &Repositories{db}
}

func (r *Repositories) RepoCreateSegments(ctx *fiber.Ctx, request *entity.ChangeSegmentsRequest, response *entity.ChangeSegmentsResponse) (*entity.ChangeSegmentsResponse, error) {
	context := ctx.Context()

	err := r.db.QueryRowContext(context, databases.CreateSegments, request.Slug).Scan(&response.Slug)
	if err != nil {
		response.Status = "Error"
		return response, nil
	}

	response.Status = "Success"
	return response, nil
}

func (r *Repositories) RepoDeleteSegments(ctx *fiber.Ctx, request *entity.ChangeSegmentsRequest, response *entity.ChangeSegmentsResponse) (*entity.ChangeSegmentsResponse, error) {
	context := ctx.Context()

	err := r.db.QueryRowContext(context, databases.DeleteSegments, request.Slug).Scan(&response.Slug)
	if err != nil {
		response.Status = "Error"
		return response, nil
	}

	response.Status = "Success"
	return response, nil
}

func (r *Repositories) RepoGetActiveUserSegments(ctx *fiber.Ctx, request *entity.CheckSegmentsRequest, response *entity.CheckSegmentsResponse) (*entity.CheckSegmentsResponse, error) {
	context := ctx.Context()

	err := r.db.QueryRowContext(context, databases.GetActiveUserSegments, request.UserID).Scan(pq.Array(&response.Segments))
	if err != nil {
		response.Status = "Error"
		return response, nil
	}

	response.Status = "Success"
	return response, nil
}
