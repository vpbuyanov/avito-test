package repos

import (
	"avito-test/internal/databases"
	"avito-test/internal/entity"
	"database/sql"
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

	err := r.db.QueryRowContext(context, databases.CreateSegments, request.Slug).Scan(&response.Slug)
	if err != nil {
		response.Status = "Error"
		return response, nil
	}

	response.Status = "Success"
	return response, nil
}

//
//func (r *Repositories) RepoInvoice(ctx *fiber.Ctx, request *entity.Request, response *entity.Response) (*entity.Response, error) {
//	context := ctx.Context()
//
//	tx, err := r.db.BeginTx(context, nil)
//	if err != nil {
//		return response, err
//	}
//
//	defer tx.Rollback()
//
//	err = tx.QueryRowContext(context, databases.FindWallet, request.NumberWallet).Scan(&response.NumberWallet, &response.СurrencyСode, &response.Amount)
//	if err != nil {
//		if errors.Is(err, sql.ErrNoRows) {
//			_, err = tx.ExecContext(context, databases.FirstDeposit, request.NumberWallet, request.СurrencyСode, request.Amount)
//			if err != nil {
//				response.Status = "Error"
//				return response, err
//			}
//
//			err = tx.QueryRowContext(context, databases.FindWallet, request.NumberWallet).Scan(&response.NumberWallet, &response.СurrencyСode, &response.Amount)
//			if err != nil {
//				response.Status = "Error"
//				return response, err
//			}
//
//			err = tx.Commit()
//			if err != nil {
//				response.Status = "Error"
//				return response, nil
//			}
//
//			response.Status = "Success"
//			return response, nil
//		}
//		return nil, err
//	}
//
//	_, err = tx.ExecContext(context, databases.Deposit, request.NumberWallet, request.СurrencyСode, request.Amount)
//	if err != nil {
//		response.Status = "Error"
//		return response, err
//	}
//
//	err = tx.QueryRowContext(context, databases.FindWallet, request.NumberWallet).Scan(&response.NumberWallet, &response.СurrencyСode, &response.Amount)
//	if err != nil {
//		response.Status = "Error"
//		return response, err
//	}
//
//	err = tx.Commit()
//	if err != nil {
//		response.Status = "Error"
//		return response, nil
//	}
//	response.Status = "Success"
//	return response, nil
//}
//
//func (r *Repositories) RepoWithdraw(ctx *fiber.Ctx, request *entity.Request, response *entity.Response) (*entity.Response, error) {
//	context := ctx.Context()
//
//	tx, err := r.db.BeginTx(context, nil)
//	if err != nil {
//		response.Status = "Error"
//		return response, err
//	}
//
//	defer tx.Rollback()
//
//	err = tx.QueryRowContext(context, databases.FindWallet, request.NumberWallet).Scan(&response.NumberWallet, &response.СurrencyСode, &response.Amount)
//	if err != nil {
//		if err == sql.ErrNoRows {
//			response.Status = "Error"
//			return response, nil
//		}
//		return response, err
//	}
//
//	if response.Amount >= request.Amount {
//		_, err = tx.ExecContext(context, databases.Withdraw, request.NumberWallet, request.СurrencyСode, request.Amount)
//		if err != nil {
//			response.Status = "Error"
//			return response, err
//		}
//
//		err = tx.QueryRowContext(context, databases.FindWallet, request.NumberWallet).Scan(&response.NumberWallet, &response.СurrencyСode, &response.Amount)
//		if err != nil {
//			if err == sql.ErrNoRows {
//				response.Status = "Error"
//				return response, nil
//			}
//			return response, err
//		}
//
//		err = tx.Commit()
//		if err != nil {
//			response.Status = "Error"
//			return response, err
//		}
//
//		response.Status = "Success"
//		return response, nil
//	}
//
//	response.Status = "Error"
//	return response, nil
//}
//
//func (r *Repositories) CheckWallet(ctx *fiber.Ctx, request *entity.CheckWalletRequest, response *entity.CheckWalletResponse) (*entity.CheckWalletResponse, error) {
//	context := ctx.Context()
//
//	tx, err := r.db.BeginTx(context, nil)
//	if err != nil {
//		response.Status = "Error"
//		return response, err
//	}
//
//	defer tx.Rollback()
//
//	err = tx.QueryRowContext(context, databases.CheckStatusWallet, request.NumberWallet).Scan(&response.NumberWallet, &response.Status)
//	if err != nil {
//		if err == sql.ErrNoRows {
//			response.Status = "Error"
//			return response, nil
//		}
//		return response, err
//	}
//
//	err = tx.Commit()
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
