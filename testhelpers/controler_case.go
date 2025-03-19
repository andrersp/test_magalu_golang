package testhelper

import (
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
)

type ControllerCase struct {
	Req     *http.Request
	Res     *httptest.ResponseRecorder
	Context echo.Context
}

// SetupControllerCase initializes a test case for the Echo controller.
func SetupControllerCase(method, url string, body io.Reader) ControllerCase {
	engine := echo.New()
	req := httptest.NewRequest(method, url, body)
	res := httptest.NewRecorder()
	ctx := engine.NewContext(req, res)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	return ControllerCase{Req: req, Res: res, Context: ctx}
}
