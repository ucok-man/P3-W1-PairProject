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

	root := router.Group("/v1")
	root.Use(app.withRecover())
	root.Use(app.withLogger())

	router.GET("/swagger/*", echoSwagger.WrapHandler)

	users := root.Group("/users")
	{
		users.POST("/register", app.userRegisterHandler)
		users.POST("/login", app.userLoginHandler)
	}

	messages := root.Group("/messages")
	messages.Use(app.withLogin)
	{
		messages.POST("", app.messageCreateHandler)
	}

	return router
}
