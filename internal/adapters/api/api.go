package api

import (
	"errors"
	"net/http"

	_ "github.com/andrersp/favorites/docs"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/andrersp/favorites/internal/adapters/api/middlewares"
	genericresponse "github.com/andrersp/favorites/pkg/generic-response"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	defaultMessage = "INTERNAL ERROR"
	defaultDetail  = "internal error on server"
)

func Api() *echo.Echo {
	server := echo.New()
	server.HTTPErrorHandler = errorHandler
	server.Use(middleware.Recover())
	server.Use(middlewares.LoggerMiddleware)
	server.GET("/swagger/*", echoSwagger.WrapHandler)

	return server
}

func errorHandler(err error, ctx echo.Context) {
	code := http.StatusBadRequest

	var genericError *genericresponse.Error

	if errors.As(err, &genericError) {
		_ = ctx.JSON(code, genericError)
		return
	}

	var echoError *echo.HTTPError

	if errors.As(err, &echoError) {
		code = echoError.Code
	}

	err = genericresponse.NewErrorResponse(defaultMessage, err.Error())

	_ = ctx.JSON(code, err)
}
