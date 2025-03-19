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

type ProductHandler struct {
	app            app.ProductApp
	authMiddleware *middlewares.AuthMiddleware
}

func NewProductHandler(
	app app.ProductApp,
	authMiddleware *middlewares.AuthMiddleware,
) *ProductHandler {
	handler := new(ProductHandler)
	handler.app = app
	handler.authMiddleware = authMiddleware

	return handler
}

func (p *ProductHandler) Setup(group *echo.Group) {
	group.GET("/products", p.List, p.authMiddleware.ValidateToken(false))
	group.GET("/products/:productID", p.Detail, p.authMiddleware.ValidateToken(false))
}

// @tags Product
// @summary list products
// @description Endpoit to list products
// @Param        page    query     int  false  "page number"
// @success 200 {array} dto.ProductResponseDTO
// @failure 400 {object} genericresponse.Error
// @Security ApiKeyAuth
// @router /products [get]
func (p *ProductHandler) List(ctx echo.Context) error {
	var queryData dto.ProductPageRequestDTO

	if err := ctx.Bind(&queryData); err != nil {
		return genericresponse.NewErrorResponse(genericresponse.INVALID_PARAM, err.Error())
	}

	if err := queryData.Validate(); err != nil {
		return genericresponse.NewErrorResponse(genericresponse.INVALID_PARAM, err.Error())
	}

	response, err := p.app.List(queryData.Page)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response)
}

// @tags Product
// @summary detail product
// @description Endpoit to detail product
// @Param        productID    path     uint  true  "product id"
// @success 200 {object} dto.ProductResponseDTO
// @failure 400 {object} genericresponse.Error
// @failure 404 {object} genericresponse.Error
// @Security ApiKeyAuth
// @router /products/{productID} [get]
func (p *ProductHandler) Detail(ctx echo.Context) error {
	var request dto.ProductDetailRequestDTO

	if err := ctx.Bind(&request); err != nil {
		return genericresponse.NewErrorResponse(genericresponse.INVALID_PARAM, err.Error())
	}

	if err := request.Validate(); err != nil {
		return genericresponse.NewErrorResponse(genericresponse.INVALID_PARAM, err.Error())
	}

	response, err := p.app.Detail(request.ID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return ctx.JSON(http.StatusNotFound,
				genericresponse.NewErrorResponse(genericresponse.NOT_FOUND, err.Error()))
		}

		return genericresponse.NewErrorResponse("error on get product detail", err.Error())
	}

	return ctx.JSON(http.StatusOK, response)
}
