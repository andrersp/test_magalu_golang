package handlers

import (
	"net/http"
	"strings"

	"github.com/andrersp/favorites/internal/app"
	"github.com/andrersp/favorites/internal/domain/dto"
	genericresponse "github.com/andrersp/favorites/pkg/generic-response"
	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	app app.ClientApp
}

func NewAdminHandler(
	app app.ClientApp,

) *AdminHandler {
	handler := new(AdminHandler)
	handler.app = app

	return handler
}

func (a *AdminHandler) Setup(group *echo.Group) {
	group.POST("/clients", a.Register)
	group.GET("/clients/:clientID", a.Detail)
	group.PUT("/clients/:clientID", a.Update)
	group.GET("/clients", a.List)
	group.DELETE("/clients/:clientID", a.Delete)
}

// @tags Admin
// @summary Client register
// @description Endpoit to register client
// @param Payload body dto.ClientRequestDTO true "Payload"
// @success 201
// @failure 400 {object} genericresponse.Error
// @Security ApiKeyAuth
// @router /admin/clients [post]
func (c *AdminHandler) Register(ctx echo.Context) error {
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

// @tags Admin
// @summary List Client
// @description Endpoit to list clients
// @success 200 {array} dto.ClientResumeResponseDTO
// @failure 400 {object} genericresponse.Error
// @Security ApiKeyAuth
// @router /admin/clients [get]
func (a *AdminHandler) List(ctx echo.Context) error {
	response, err := a.app.List()
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response)
}

// @tags Admin
// @summary Client Detail
// @description Endpoit client detail
// @param clientID path uint true "Client id" example(1)
// @success 200 {object} dto.ClientDetailResponseDTO
// @failure 400 {object} genericresponse.Error
// @failure 404 {object} genericresponse.Error
// @Security ApiKeyAuth
// @router /admin/clients/{clientID} [get]
func (a *AdminHandler) Detail(ctx echo.Context) error {
	var request dto.ClientDetailRequestDTO

	if err := ctx.Bind(&request); err != nil {
		return genericresponse.NewErrorResponse(genericresponse.INVALID_PARAM, err.Error())
	}

	if err := request.Validate(); err != nil {
		return genericresponse.NewErrorResponse(genericresponse.INVALID_PARAM, err.Error())
	}

	response, err := a.app.Detail(request.ID)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return ctx.JSON(http.StatusNotFound,
				genericresponse.NewErrorResponse(genericresponse.NOT_FOUND, err.Error()))
		}

		return genericresponse.NewErrorResponse("error on get client detail", err.Error())
	}

	return ctx.JSON(http.StatusOK, response)
}

// @tags Admin
// @summary Client Update
// @description Endpoit client update
// @param clientID path uint true "Client id" example(1)
// @param Payload body dto.ClientRequestDTO true "Payload"
// @success 204
// @failure 400 {object} genericresponse.Error
// @failure 404 {object} genericresponse.Error
// @Security ApiKeyAuth
// @router /admin/clients/{clientID} [put]
func (a *AdminHandler) Update(ctx echo.Context) error {
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

	err := a.app.Update(requestParam.ID, payload)

	if err != nil {
		return genericresponse.NewErrorResponse("error on get client detail", err.Error())
	}

	return ctx.NoContent(http.StatusNoContent)
}

// @tags Admin
// @summary Client Delete
// @description Endpoit client delete
// @param clientID path uint true "Client id" example(1)
// @success 204
// @failure 400 {object} genericresponse.Error
// @Security ApiKeyAuth
// @router /clients/{clientID} [delete]
func (a *AdminHandler) Delete(ctx echo.Context) error {
	var requestParam dto.ClientDetailRequestDTO

	if err := (&echo.DefaultBinder{}).BindPathParams(ctx, &requestParam); err != nil {
		return genericresponse.NewErrorResponse(genericresponse.INVALID_PARAM, err.Error())
	}

	if err := requestParam.Validate(); err != nil {
		return genericresponse.NewErrorResponse(genericresponse.INVALID_PARAM, err.Error())
	}

	err := a.app.Delete(requestParam.ID)

	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}
