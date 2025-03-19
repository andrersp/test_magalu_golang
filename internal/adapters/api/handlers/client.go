package handlers

import (
	"net/http"
	"strings"

	"github.com/andrersp/favorites/internal/adapters/api/middlewares"
	"github.com/andrersp/favorites/internal/app"
	"github.com/andrersp/favorites/internal/domain/dto"
	genericresponse "github.com/andrersp/favorites/pkg/generic-response"
	"github.com/labstack/echo/v4"
)

type ClientHandler struct {
	app            app.ClientApp
	authMiddleware *middlewares.AuthMiddleware
}

func NewClientHandler(
	app app.ClientApp,
	authMiddleware *middlewares.AuthMiddleware,
) *ClientHandler {
	handler := new(ClientHandler)
	handler.app = app
	handler.authMiddleware = authMiddleware

	return handler
}

func (a *ClientHandler) Setup(group *echo.Group) {
	group.POST("/clients", a.Register)
	group.GET("/clients/:clientID", a.Detail, a.authMiddleware.ValidateToken(false))
	group.PUT("/clients/:clientID", a.Update, a.authMiddleware.ValidateToken(false))
}

// @tags Client
// @summary Client register
// @description Endpoit to self register
// @param Payload body dto.ClientRequestDTO true "Payload"
// @success 201
// @failure 400 {object} genericresponse.Error
// @router /clients [post]
func (c *ClientHandler) Register(ctx echo.Context) error {
	var payload dto.ClientRequestDTO

	if err := ctx.Bind(&payload); err != nil {
		return genericresponse.NewErrorResponse(genericresponse.INVALID_PAYLOAD, err.Error())
	}

	if err := payload.Validate(); err != nil {
		return genericresponse.NewErrorResponse(genericresponse.INVALID_PAYLOAD, err.Error())
	}

	err := c.app.Create(payload)

	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusCreated)
}

// @tags Client
// @summary Client Detail
// @description Endpoit client detail
// @param clientID path uint true "Client id" example(1)
// @success 200 {object} dto.ClientDetailResponseDTO
// @failure 400 {object} genericresponse.Error
// @failure 404 {object} genericresponse.Error
// @Security ApiKeyAuth
// @router /clients/{clientID} [get]
func (c *ClientHandler) Detail(ctx echo.Context) error {
	var request dto.ClientDetailRequestDTO

	if err := ctx.Bind(&request); err != nil {
		return genericresponse.NewErrorResponse(genericresponse.INVALID_PARAM, err.Error())
	}

	if err := request.Validate(); err != nil {
		return genericresponse.NewErrorResponse(genericresponse.INVALID_PARAM, err.Error())
	}

	response, err := c.app.Detail(request.ID)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return ctx.JSON(http.StatusNotFound,
				genericresponse.NewErrorResponse(genericresponse.NOT_FOUND, err.Error()))
		}

		return genericresponse.NewErrorResponse("error on get client detail", err.Error())
	}

	return ctx.JSON(http.StatusOK, response)
}

// @tags Client
// @summary Client Update
// @description Endpoit client update
// @param clientID path uint true "Client id" example(1)
// @param Payload body dto.ClientRequestDTO true "Payload"
// @success 204
// @failure 400 {object} genericresponse.Error
// @failure 404 {object} genericresponse.Error
// @Security ApiKeyAuth
// @router /clients/{clientID} [put]
func (c *ClientHandler) Update(ctx echo.Context) error {
	var (
		requestParam dto.ClientDetailRequestDTO
		payload      dto.ClientRequestDTO
	)

	if err := (&echo.DefaultBinder{}).BindPathParams(ctx, &requestParam); err != nil {
		return genericresponse.NewErrorResponse(genericresponse.INVALID_PARAM, err.Error())
	}

	if err := requestParam.Validate(); err != nil {
		return genericresponse.NewErrorResponse(genericresponse.INVALID_PARAM, err.Error())
	}

	if err := (&echo.DefaultBinder{}).BindBody(ctx, &payload); err != nil {
		return genericresponse.NewErrorResponse(genericresponse.INVALID_PAYLOAD, err.Error())
	}

	if err := payload.Validate(); err != nil {
		return genericresponse.NewErrorResponse(genericresponse.INVALID_PAYLOAD, err.Error())
	}

	err := c.app.Update(requestParam.ID, payload)

	if err != nil {
		return genericresponse.NewErrorResponse("error on get client detail", err.Error())
	}

	return ctx.NoContent(http.StatusNoContent)
}
