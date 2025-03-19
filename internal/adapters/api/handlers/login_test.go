package handlers_test

import (
	"net/http"
	"strings"
	"testing"

	"github.com/andrersp/favorites/internal/adapters/api/handlers"
	"github.com/andrersp/favorites/internal/adapters/repository/client"
	"github.com/andrersp/favorites/internal/adapters/security"
	"github.com/andrersp/favorites/internal/app"
	"github.com/andrersp/favorites/internal/domain/dto"
	testhelper "github.com/andrersp/favorites/testhelpers"

	"github.com/stretchr/testify/suite"
)

const (
	tokenSecret     = "hRjU/cU1TJJiogiFQud5+/bkTFmkehfH"
	tokenExpiration = 10
	loginURL        = "/login"
)

var (
	loginPayloadValid = dto.LoginRequestDTO{
		Email: testhelper.DEFAULT_CLIENT_EMAIL,
	}
	loginPayloadInvalid = dto.LoginRequestDTO{
		Email: "admin@mail.com",
	}

	loginPayloadAdmin = dto.LoginRequestDTO{
		Email: "admin@admin.com",
	}
)

type loginTest struct {
	suite.Suite
	handler *handlers.LoginHandler
}

func TestLoginHandler(t *testing.T) {
	suite.Run(t, new(loginTest))
}

func (suite *loginTest) SetupSuite() {
	t := suite.T()
	containerW := testhelper.NewPostgres(t)
	gorm, err := testhelper.GetGormInstance(t, containerW)

	suite.NoError(err)

	clientRepo := client.NewClientRepository(gorm.DB())
	tokenService := security.NewTokenService(tokenSecret, tokenExpiration)

	loginApp := app.NewLoginApp(clientRepo, tokenService)

	suite.handler = handlers.NewLoginHandler(loginApp)
}

func (suite *loginTest) TestLoginErrorOnBindPayload() {
	payload := `{"email": 0}`

	req := testhelper.SetupControllerCase(http.MethodPost, loginURL, strings.NewReader(payload))

	err := suite.handler.Login(req.Context)

	suite.Error(err)
}

func (suite *loginTest) TestLoginErrorOnValidatePayload() {
	payload := `{"email": "mail@mail"}`

	req := testhelper.SetupControllerCase(http.MethodPost, loginURL, strings.NewReader(payload))

	err := suite.handler.Login(req.Context)

	suite.Error(err)
}

func (suite *loginTest) TestLoginSuccessLogin() {
	payload := testhelper.RequestToPayload(loginPayloadValid)

	req := testhelper.SetupControllerCase(http.MethodPost, loginURL, payload)

	err := suite.handler.Login(req.Context)

	suite.NoError(err)
}

func (suite *loginTest) TestLoginErrorLogin() {
	payload := testhelper.RequestToPayload(loginPayloadInvalid)

	req := testhelper.SetupControllerCase(http.MethodPost, loginURL, payload)

	_ = suite.handler.Login(req.Context)
	suite.Equal(req.Res.Code, http.StatusBadRequest)
}

func (suite *loginTest) TestLoginSuccessLoginAdmin() {
	payload := testhelper.RequestToPayload(loginPayloadAdmin)

	req := testhelper.SetupControllerCase(http.MethodPost, loginURL, payload)

	_ = suite.handler.Login(req.Context)
	suite.Equal(req.Res.Code, http.StatusOK)
}
