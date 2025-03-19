package handlers_test

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/andrersp/favorites/internal/adapters/api/handlers"
	"github.com/andrersp/favorites/internal/adapters/api/middlewares"
	"github.com/andrersp/favorites/internal/adapters/repository/cache"
	"github.com/andrersp/favorites/internal/adapters/repository/product"
	"github.com/andrersp/favorites/internal/adapters/security"
	"github.com/andrersp/favorites/internal/app"
	testhelper "github.com/andrersp/favorites/testhelpers"

	"github.com/stretchr/testify/suite"
)

const (
	productURL = "/products"
)

type productTest struct {
	suite.Suite
	handler *handlers.ProductHandler
}

func TestProductHandler(t *testing.T) {
	suite.Run(t, new(productTest))
}

func (suite *productTest) SetupSuite() {
	productRepository := product.NewFakeProductRepository()
	cacheRepository := cache.NewCacheRepository()
	tokenService := security.NewTokenService(tokenSecret, tokenExpiration)
	authMiddleware := middlewares.NewAuthMiddleware(tokenService)

	productApp := app.NewProductApp(productRepository, cacheRepository)

	suite.handler = handlers.NewProductHandler(productApp, authMiddleware)
}

func (suite *productTest) TestProductErrorOnBindPayload() {
	q := make(url.Values)
	q.Set("page", "1aa")

	req := testhelper.SetupControllerCase(http.MethodGet, fmt.Sprintf("%s?%s", productURL, q.Encode()), nil)

	err := suite.handler.List(req.Context)

	suite.Error(err)
}

func (suite *productTest) TestProductErrorOnValidatePayload() {
	q := make(url.Values)
	q.Set("page", "0")

	req := testhelper.SetupControllerCase(http.MethodGet, fmt.Sprintf("%s?%s", productURL, q.Encode()), nil)

	err := suite.handler.List(req.Context)

	suite.Error(err)
}

func (suite *productTest) TestProductListSuccess() {
	q := make(url.Values)
	q.Set("page", "1")

	req := testhelper.SetupControllerCase(http.MethodGet, fmt.Sprintf("%s?%s", productURL, q.Encode()), nil)

	err := suite.handler.List(req.Context)
	suite.NoError(err)
}

func (suite *productTest) TestProductDetailErrorBind() {
	req := testhelper.SetupControllerCase(http.MethodGet, productURL, nil)
	req.Context.SetParamNames("productId")
	req.Context.SetParamValues("1a")

	err := suite.handler.Detail(req.Context)
	suite.Error(err)
}

func (suite *productTest) TestProductDetailErrorvalidate() {
	req := testhelper.SetupControllerCase(http.MethodGet, productURL, nil)
	req.Context.SetParamNames("productId")
	req.Context.SetParamValues("0")

	err := suite.handler.Detail(req.Context)
	suite.Error(err)
}

func (suite *productTest) TestProductDetailSuccess() {
	req := testhelper.SetupControllerCase(http.MethodGet, productURL, nil)
	req.Context.SetParamNames("productId")
	req.Context.SetParamValues("1")

	err := suite.handler.Detail(req.Context)
	suite.NoError(err)
}

func (suite *productTest) TestProductDetailErrorProductNotFound() {
	req := testhelper.SetupControllerCase(http.MethodGet, productURL, nil)
	req.Context.SetParamNames("productId")
	req.Context.SetParamValues("1111")

	err := suite.handler.Detail(req.Context)
	if suite.NoError(err) {
		suite.Equal(http.StatusNotFound, req.Res.Code)
	}
}
