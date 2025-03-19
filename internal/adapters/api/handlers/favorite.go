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

type FavoriteHandler struct {
	app            app.FavoriteApp
	authMiddleware *middlewares.AuthMiddleware
}

func NewFavoriteHandler(
	app app.FavoriteApp,
	authMiddleware *middlewares.AuthMiddleware,
) *FavoriteHandler {
	handler := new(FavoriteHandler)
	handler.app = app
	handler.authMiddleware = authMiddleware

	return handler
}

func (f *FavoriteHandler) Setup(group *echo.Group) {
	group.POST("/favorites", f.Add, f.authMiddleware.ValidateToken(false))
	group.DELETE("/favorites/:clientId/:productId", f.Delete, f.authMiddleware.ValidateToken(false))
}

// @tags Favorite
// @summary Add Favorite
// @description Endpoit add favorite
// @param Payload body dto.FavoriteRequestDTO true "Payload"
// @success 201
// @failure 400 {object} genericresponse.Error
// @Security ApiKeyAuth
// @router /favorites [post]
func (f *FavoriteHandler) Add(ctx echo.Context) error {
	var paylaod dto.FavoriteRequestDTO

	if err := ctx.Bind(&paylaod); err != nil {
		return genericresponse.NewErrorResponse(genericresponse.INVALID_PAYLOAD, err.Error())
	}

	if err := paylaod.Validate(); err != nil {
		return genericresponse.NewErrorResponse(genericresponse.INVALID_PAYLOAD, err.Error())
	}

	err := f.app.Add(paylaod)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return ctx.JSON(http.StatusNotFound,
				genericresponse.NewErrorResponse(genericresponse.NOT_FOUND, err.Error()))
		}

		return err
	}

	return ctx.NoContent(http.StatusCreated)
}

// @tags Favorite
// @summary Delete Favorite
// @description Endpoit delete favorite
// @Param        clientId    path     uint  true  "client id"
// @Param        productId    path     uint  true  "product id"
// @success 204
// @failure 400 {object} genericresponse.Error
// @Security ApiKeyAuth
// @router /favorites/{clientId}/{productId} [delete]
func (f *FavoriteHandler) Delete(ctx echo.Context) error {
	var paylaod dto.FavoriteRequestDTO

	if err := ctx.Bind(&paylaod); err != nil {
		return genericresponse.NewErrorResponse(genericresponse.INVALID_PAYLOAD, err.Error())
	}

	if err := paylaod.Validate(); err != nil {
		return genericresponse.NewErrorResponse(genericresponse.INVALID_PAYLOAD, err.Error())
	}

	err := f.app.Delete(paylaod)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return ctx.JSON(http.StatusNotFound,
				genericresponse.NewErrorResponse(genericresponse.NOT_FOUND, err.Error()))
		}

		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}
