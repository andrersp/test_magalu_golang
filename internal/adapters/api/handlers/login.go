package handlers

import (
	"net/http"

	"github.com/andrersp/favorites/internal/app"
	"github.com/andrersp/favorites/internal/domain/dto"
	genericresponse "github.com/andrersp/favorites/pkg/generic-response"
	"github.com/labstack/echo/v4"
)

type LoginHandler struct {
	loginApp app.LoginApp
}

func NewLoginHandler(
	loginApp app.LoginApp,
) *LoginHandler {
	handler := new(LoginHandler)
	handler.loginApp = loginApp

	return handler
}

// @tags Login
// @summary Login
// @description Endpoit login
// @param Payload body dto.LoginRequestDTO true "Payload"
// @success 200 {object} dto.LoginResponseDTO
// @failure 400 {object} genericresponse.Error
// @router /login [post]
func (l *LoginHandler) Setup(group *echo.Group) {
	group.POST("/login", l.Login)
}

// @tags Login
// @summary Login
// @description Endpoit login
// @param Payload body dto.LoginRequestDTO true "Payload"
// @success 200 {object} dto.LoginResponseDTO
// @failure 400 {object} genericresponse.Error
// @router /login [post]
func (l *LoginHandler) Login(ctx echo.Context) error {
	var payload dto.LoginRequestDTO

	if err := ctx.Bind(&payload); err != nil {
		return genericresponse.NewErrorResponse(genericresponse.INVALID_PAYLOAD, err.Error())
	}

	if err := payload.Validate(); err != nil {
		return genericresponse.NewErrorResponse(genericresponse.INVALID_PAYLOAD, err.Error())
	}

	response, err := l.loginApp.Login(payload.Email)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, genericresponse.NewErrorResponse(genericresponse.INVALID_PAYLOAD, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response)
}
