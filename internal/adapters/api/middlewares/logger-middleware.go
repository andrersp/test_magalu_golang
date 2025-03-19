package middlewares

import (
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

func LoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		start := time.Now()

		if err := next(ctx); err != nil {
			ctx.Error(err)
		}

		if strings.Contains(ctx.Request().RequestURI, "swagger") {
			return nil
		}

		duration := time.Since(start).Seconds()
		uri := ctx.Request().URL.Path
		method := ctx.Request().Method
		statusCode := ctx.Response().Status
		message := []any{"method", method, "path", uri, "duration", duration, "status", statusCode}

		if statusCode < http.StatusBadRequest {
			slog.Info("request", message...)
			return nil
		}

		slog.Error("request", message...)

		return nil
	}
}
