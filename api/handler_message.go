package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/ucok-man/P3-W1-PairProject/internal/contract"
)

func (app *Application) messageCreateHandler(ctx echo.Context) error {
	var input contract.ReqMessageCreate

	if err := ctx.Bind(&input); err != nil {
		return app.ErrBadRequest(ctx, err)
	}

	if err := ctx.Validate(&input); err != nil {
		return app.ErrFailedValidation(ctx, err)
	}

	// var chatbox
	return app.ErrInternalServer(ctx, fmt.Errorf("UNIMPLEMENTED ENDPOINT!"))
}
