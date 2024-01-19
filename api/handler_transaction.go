package api

import (
	"errors"
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/ucok-man/P3-W1-PairProject/internal/contract"
	"github.com/ucok-man/P3-W1-PairProject/internal/entity"
	"github.com/ucok-man/P3-W1-PairProject/internal/repo"
)

// TODO: yang get masuk sini gan: oke gan

func (app *Application) transactionInsertHandler(ctx echo.Context) error {
	var input contract.ReqTransactionInsert
	if err := ctx.Bind(&input); err != nil {
		return app.ErrBadRequest(ctx, err)
	}

	if err := ctx.Validate(&input); err != nil {
		return app.ErrFailedValidation(ctx, err)
	}

	var transaction = &entity.Transaction{}
	if err := copier.Copy(&transaction, &input); err != nil {
		return app.ErrInternalServer(ctx, err)
	}

	err := app.repo.Transaction.Insert(ctx.Request().Context(), transaction)
	if err != nil {
		return app.ErrInternalServer(ctx, err)
	}

	var response contract.ResTransactionInsert
	if err := copier.Copy(&response.Data, transaction); err != nil {
		return app.ErrInternalServer(ctx, err)
	}
	response.Message = "success"

	return ctx.JSON(http.StatusCreated, response)
}

func (app *Application) transactionGetAllHandler(ctx echo.Context) error {
	transactions, err := app.repo.Transaction.GetAll(ctx.Request().Context())
	if err != nil {
		return app.ErrInternalServer(ctx, err)
	}

	var response contract.ResTransactionGetAll
	if err := copier.Copy(&response.Data, transactions); err != nil {
		return app.ErrInternalServer(ctx, err)
	}

	return ctx.JSON(http.StatusCreated, response)
}

func (app *Application) transactionGetByIdHandler(ctx echo.Context) error {
	transactionID, err := app.getParamId(ctx)
	if err != nil {
		return app.ErrBadRequest(ctx, err)
	}

	transaction, err := app.repo.Transaction.GetByID(ctx.Request().Context(), transactionID)
	if err != nil {
		switch {
		case errors.Is(err, repo.ErrRecordNotFound):
			return app.ErrNotFound(ctx)
		default:
			return app.ErrInternalServer(ctx, err)
		}
	}

	var response contract.ResTransactionGetByID
	if err := copier.Copy(&response.Data, transaction); err != nil {
		return app.ErrInternalServer(ctx, err)
	}

	return ctx.JSON(http.StatusCreated, response)
}

func (app *Application) transactionUpdateHandler(ctx echo.Context) error {
	var input contract.ReqTransactionUpdate
	if err := ctx.Bind(&input); err != nil {
		return app.ErrBadRequest(ctx, err)
	}

	if err := ctx.Validate(&input); err != nil {
		return app.ErrFailedValidation(ctx, err)
	}

	var transaction entity.Transaction
	if err := copier.Copy(&transaction, &input); err != nil {
		return app.ErrInternalServer(ctx, err)
	}

	res, err := app.repo.Transaction.Update(ctx.Request().Context(), transaction)
	if err != nil {
		return app.ErrInternalServer(ctx, err)
	}

	return ctx.JSON(http.StatusOK, res)
}

func (app *Application) transactionDeleteHandler(ctx echo.Context) error {
	var input contract.ReqTransactionDelete
	if err := ctx.Bind(&input); err != nil {
		return app.ErrBadRequest(ctx, err)
	}

	if err := ctx.Validate(&input); err != nil {
		return app.ErrFailedValidation(ctx, err)
	}

	var transaction entity.Transaction
	if err := copier.Copy(&transaction, &input); err != nil {
		return app.ErrInternalServer(ctx, err)
	}

	res, err := app.repo.Transaction.Delete(ctx.Request().Context(), transaction)
	if err != nil {
		return app.ErrInternalServer(ctx, err)
	}

	return ctx.JSON(http.StatusOK, res)
}

func (app *Application) transactionDeleteAllHandler(ctx echo.Context) error {
	res, err := app.repo.Transaction.DeleteAll(ctx.Request().Context())
	if err != nil {
		return app.ErrInternalServer(ctx, err)
	}

	return ctx.JSON(http.StatusOK, res)
}
