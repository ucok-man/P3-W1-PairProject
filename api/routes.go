package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/ucok-man/P3-W1-PairProject/internal/serializer"
	"github.com/ucok-man/P3-W1-PairProject/internal/validator"
)

func (app *Application) routes() http.Handler {
	router := echo.New()
	router.HTTPErrorHandler = app.httpErrorHandler
	router.JSONSerializer = serializer.JSONSerializer{}
	router.Validator = validator.New()

	router.GET("/swagger/*", echoSwagger.WrapHandler)

	root := router.Group("/v1")
	root.Use(app.withRecover())
	root.Use(app.withLogger())

	// users := root.Group("/users")
	// {
	// 	users.POST("/register", app.userRegisterHandler)
	// 	users.POST("/login", app.userLoginHandler)
	// }

	transactions := root.Group("/transactions")
	{
		// TODO: yang get masuk sini gan: oke gan
		transactions.POST("", app.transactionInsertHandler)
		transactions.GET("", app.transactionGetAllHandler)
		transactions.GET("/:id", app.transactionGetByIdHandler)
		transactions.PUT("/:id", app.transactionUpdateHandler)
		transactions.DELETE("/:id", app.transactionDeleteHandler)
	}

	return router
}
