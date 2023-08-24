package usecase

import (
	"avito-test/internal/databases/repos"
	"avito-test/internal/entity"
	"github.com/gofiber/fiber/v2"
)

type UseCase struct {
	repo Repo
}

func NewUseCase(r *repos.Repositories) *UseCase {
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

//func (use *UseCase) UseCaseInvoice(ctx *fiber.Ctx, request *entity.Request, response *entity.Response) (*entity.Response, error) {
//	response, err := use.repo.RepoInvoice(ctx, request, response)
//	if err != nil {
//		response.Status = "Error"
//		return response, err
//	}
//
//	return response, nil
//}
//
//func (use *UseCase) UseCaseWithdraw(ctx *fiber.Ctx, request *entity.Request, response *entity.Response) (*entity.Response, error) {
//	response, err := use.repo.RepoWithdraw(ctx, request, response)
//	if err != nil {
//		response.Status = "Error"
//		return response, nil
//	}
//	return response, nil
//}
//
//func (use *UseCase) UseCaseCheckStatusWallet(ctx *fiber.Ctx, request *entity.CheckWalletRequest, response *entity.CheckWalletResponse) (*entity.CheckWalletResponse, error) {
//	response, err := use.repo.CheckWallet(ctx, request, response)
//	if err != nil {
//		response.Status = "Error"
//		return response, nil
//	}
//	return response, nil
//}
