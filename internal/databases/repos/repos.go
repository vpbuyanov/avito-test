package repos

import (
	"avito-test/internal/databases"
	"avito-test/internal/entity"
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
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

	err := r.db.QueryRowContext(context, databases.CreateSegments, request.Slug).Scan(&response.SegmentsId, &response.Slug)
	if err != nil {
		response.Status = "Error"
		return response, nil
	}

	response.Status = "Success"
	return response, nil
}

func (r *Repositories) RepoDeleteSegments(ctx *fiber.Ctx, request *entity.ChangeSegmentsRequest, response *entity.ChangeSegmentsResponse) (*entity.ChangeSegmentsResponse, error) {
	context := ctx.Context()

	err := r.db.QueryRowContext(context, databases.DeleteSegments, request.Slug).Scan(&response.SegmentsId, &response.Slug)
	if err != nil {
		response.Status = "Error"
		return response, nil
	}

	response.Status = "Success"
	return response, nil
}

func (r *Repositories) RepoCreateUser(ctx *fiber.Ctx, request *entity.CreateUserRequest, response *entity.CreateUserResponse) (*entity.CreateUserResponse, error) {
	context := ctx.Context()

	err := r.db.QueryRowContext(context, databases.CreateUser, request.Name).Scan(&response.UserID)
	if err != nil {
		response.Status = "Error"
		return response, nil
	}

	response.Status = "Success"
	return response, nil
}

func (r *Repositories) RepoSetUserSegments(ctx *fiber.Ctx, request *entity.ChangeUserSegmentsRequest, response *entity.ChangeUserSegmentsResponse) (*entity.ChangeUserSegmentsResponse, error) {
	context := ctx.Context()

	if request.AddSlug != nil {
		for i := 0; i < len(request.AddSlug); i++ {
			var segment string
			err := r.db.QueryRowContext(context, databases.SetUserSegments, request.UserID, request.AddSlug[i]).Scan(&response.UserID, &segment)
			if err != nil {
				fmt.Println("1", err)
				response.Status = "Error"
				return response, nil
			}

			response.Segments = append(response.Segments, segment)
		}
	}

	if request.DelSlug != nil {
		for i := 0; i < len(request.DelSlug); i++ {
			var segment string
			err := r.db.QueryRowContext(context, databases.DeleteUserSegments, request.UserID, request.DelSlug[i]).Scan(&response.UserID, &segment)
			if err != nil {
				fmt.Println("2", err)
				response.Status = "Error"
				return response, nil
			}
			response.Segments = append(response.Segments, segment)
		}
	}

	response.Status = "Success"
	return response, nil
}

func (r *Repositories) RepoGetActiveUserSegments(ctx *fiber.Ctx, request *entity.CheckSegmentsRequest, response *entity.CheckSegmentsResponse) (*entity.CheckSegmentsResponse, error) {
	context := ctx.Context()

	rows, err := r.db.QueryContext(context, databases.GetActiveUserSegments, request.UserID)
	if err != nil {
		response.Status = "Error"
		return response, nil
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			response.Status = "Error"
			return
		}
	}(rows)

	for rows.Next() {
		var segment string
		if err = rows.Scan(&segment); err != nil {
			response.Status = "Error"
			return response, nil
		}
		response.Segments = append(response.Segments, segment)
	}

	if err = rows.Err(); err != nil {
		response.Status = "Error"
		return response, nil
	}

	response.Status = "Success"
	return response, nil
}
