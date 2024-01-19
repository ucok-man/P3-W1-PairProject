package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// func (app *Application) getCurrentUser(ctx echo.Context) *entity.User {
// 	obj := ctx.Get(app.ctxkey.user)
// 	user, ok := obj.(*entity.User)
// 	if !ok {
// 		panic("[app.getCurrentUser]: user should be *entity.User")
// 	}
// 	return user
// }

func (app *Application) getParamId(ctx echo.Context) (primitive.ObjectID, error) {
	objid, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return objid, nil
}

func (app *Application) background(fn func()) {
	app.wg.Add(1)

	go func() {

		defer app.wg.Done()

		defer func() {
			if err := recover(); err != nil {
				app.logger.Error(fmt.Errorf("%s", err), "failed processing background task", nil)
			}
		}()

		fn()
	}()
}
