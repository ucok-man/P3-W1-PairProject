package api

import (
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/ucok-man/P3-W1-PairProject/internal/contract"
	"github.com/ucok-man/P3-W1-PairProject/internal/entity"
)

// TODO: yang get masuk sini gan

func (app *Application) transactionCreateHandler(ctx echo.Context) error {
	var input contract.ReqTransactionCreate
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

	res, err := app.repo.Transaction.Create(ctx.Request().Context(), transaction)
	if err != nil {
		return app.ErrInternalServer(ctx, err)
	}

	return ctx.JSON(http.StatusOK, res)
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
