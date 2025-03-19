package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/andrersp/favorites/config"
	"github.com/andrersp/favorites/internal/adapters/api"
	"github.com/andrersp/favorites/internal/adapters/api/handlers"
	"github.com/andrersp/favorites/internal/adapters/api/middlewares"
	"github.com/andrersp/favorites/internal/adapters/repository/cache"
	"github.com/andrersp/favorites/internal/adapters/repository/client"
	"github.com/andrersp/favorites/internal/adapters/repository/favorite"
	"github.com/andrersp/favorites/internal/adapters/repository/product"
	"github.com/andrersp/favorites/internal/adapters/resources/gorm"
	"github.com/andrersp/favorites/internal/adapters/resources/migration"
	"github.com/andrersp/favorites/internal/adapters/security"
	"github.com/andrersp/favorites/internal/app"
	"github.com/labstack/echo/v4"
)

func init() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug,
		AddSource: true,
	}))
	slog.SetDefault(logger)

	if err := config.LoadConfig(); err != nil {
		log.Fatal(err)
	}
}

// @title Favorites Api
// @version 1.0
// @description Api to manage clients favorites products

// @accept json
// @produce json

// @contact.name API Support
// @contact.url https://www.linkedin.com/in/rspandre/
// @contact.email rsp.assistencia@gmail.com

// @securityDefinitions.apikey  ApiKeyAuth
// @in header
// @name Authorization

// @BasePath /v1
func main() {
	dbConfig := config.GetBDConfig()

	gorm := gorm.Connect(&gorm.PgOptions{
		Host:     dbConfig.Host,
		User:     dbConfig.User,
		Password: dbConfig.Password,
		DBName:   dbConfig.Name,
		Port:     dbConfig.Port,
	}, gorm.PgConfig{})

	err := migration.Migrate("./db/migrations", gorm.URL())
	if err != nil {
		log.Fatal(err)
	}

	apiConfig := config.GetApiConfig()

	api := api.Api()

	// v1 group
	v1 := api.Group("/v1")

	// v1 Admin group
	v1Admin := v1.Group("/admin")
	SetupV1Admin(v1Admin, gorm)

	SetupV1(v1, gorm)

	if err := api.Start(fmt.Sprintf(":%s", apiConfig.Port)); err != nil {
		log.Fatal(err)
	}
}

func SetupV1(group *echo.Group, gorm *gorm.GormInstance) {
	// repository
	clientRepo := client.NewClientRepository(gorm.DB())
	productRepository := product.NewFakeProductRepository()
	cacheRepository := cache.NewCacheRepository()
	favoriteRepository := favorite.NewFavoriteRepository(gorm.DB())

	tokenConfig := config.GetConfig().Token

	tokenService := security.NewTokenService(tokenConfig.Secret, tokenConfig.Expiration)

	// App
	clientApp := app.NewClientApp(clientRepo, productRepository, cacheRepository)
	productApp := app.NewProductApp(productRepository, cacheRepository)
	loginApp := app.NewLoginApp(clientRepo, tokenService)

	favoriteApp := app.NewFavoriteApp(
		favoriteRepository, productRepository,
		clientRepo, cacheRepository,
	)

	authMiddleware := middlewares.NewAuthMiddleware(tokenService)

	// handlers
	loginHandler := handlers.NewLoginHandler(loginApp)
	loginHandler.Setup(group)

	clientHandler := handlers.NewClientHandler(clientApp, authMiddleware)
	clientHandler.Setup(group)

	productHandler := handlers.NewProductHandler(productApp, authMiddleware)
	productHandler.Setup(group)

	favoriteHandler := handlers.NewFavoriteHandler(favoriteApp, authMiddleware)
	favoriteHandler.Setup(group)
}

func SetupV1Admin(group *echo.Group, gorm *gorm.GormInstance) {
	// repository
	clientRepo := client.NewClientRepository(gorm.DB())
	productRepository := product.NewFakeProductRepository()
	cacheRepository := cache.NewCacheRepository()

	tokenConfig := config.GetConfig().Token

	tokenService := security.NewTokenService(tokenConfig.Secret, tokenConfig.Expiration)

	// App
	clientApp := app.NewClientApp(clientRepo, productRepository, cacheRepository)

	authMiddleware := middlewares.NewAuthMiddleware(tokenService)

	group.Use(authMiddleware.ValidateToken(true))

	// handlers
	adminHandler := handlers.NewAdminHandler(clientApp)
	adminHandler.Setup(group)
}
