package middlewares

import (
	"net/http"
	"strings"

	"github.com/andrersp/favorites/internal/domain/entity"
	"github.com/andrersp/favorites/internal/domain/ports"
	genericresponse "github.com/andrersp/favorites/pkg/generic-response"
	"github.com/labstack/echo/v4"
)

const (
	lengthBearer = 2
)

type AuthMiddleware struct {
	securityService ports.Security
}

func NewAuthMiddleware(securityService ports.Security) *AuthMiddleware {
	return &AuthMiddleware{
		securityService: securityService,
	}
}

func (a *AuthMiddleware) ValidateToken(isAdmin bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			token, err := extractToken(ctx.Request())
			if err != nil {
				return ctx.JSON(http.StatusUnauthorized,
					genericresponse.NewErrorResponse(genericresponse.INVALID_TOKEN, err.Error()))
			}

			tokenData, err := a.securityService.ValidateToken(token)
			if err != nil {
				return ctx.JSON(http.StatusUnauthorized,
					genericresponse.NewErrorResponse(genericresponse.INVALID_TOKEN, err.Error()))
			}

			if isAdmin && tokenData.Role != entity.ADMIN {
				return ctx.JSON(http.StatusForbidden, genericresponse.NewErrorResponse(genericresponse.FORBIDDEN, "forbidden"))
			}

			ctx.Set(entity.SecurityDataKey, tokenData)

			return next(ctx)
		}
	}
}

func extractToken(request *http.Request) (string, error) {
	authorization := request.Header.Get("Authorization")
	if authorization == "" {
		return "", genericresponse.NewErrorResponse(genericresponse.INVALID_TOKEN, "invalid token")
	}

	stringTokenList := strings.Split(authorization, " ")

	if len(stringTokenList) < lengthBearer {
		return "", genericresponse.NewErrorResponse(genericresponse.INVALID_TOKEN, "invalid token")
	}

	token := stringTokenList[1]

	return token, nil
}
