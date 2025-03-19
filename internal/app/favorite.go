package app

import (
	"fmt"
	"time"

	"github.com/andrersp/favorites/internal/domain/dto"
	"github.com/andrersp/favorites/internal/domain/ports"
)

type FavoriteApp interface {
	Add(request dto.FavoriteRequestDTO) error
	Delete(request dto.FavoriteRequestDTO) error
}

type favoriteApp struct {
	favoriteRepo ports.FavoriteRepository
	productRepo  ports.ProductRepository
	clientRepo   ports.ClientRepository
	cacheRepo    ports.CacheRepository
}

func NewFavoriteApp(
	favoriteRepo ports.FavoriteRepository,
	productRepo ports.ProductRepository,
	clientRepo ports.ClientRepository,
	cacheRepo ports.CacheRepository,
) FavoriteApp {
	app := new(favoriteApp)
	app.clientRepo = clientRepo
	app.productRepo = productRepo
	app.favoriteRepo = favoriteRepo
	app.cacheRepo = cacheRepo

	return app
}

// Add implements FavoriteApp.
func (f *favoriteApp) Add(request dto.FavoriteRequestDTO) error {
	if err := f.getClientByID(request.ClientID); err != nil {
		return err
	}

	if err := f.getProductByID(request.ProductID); err != nil {
		return err
	}

	favoriteEntity := request.ToEntity()

	_ = f.favoriteRepo.Delete(favoriteEntity)

	return f.favoriteRepo.Add(favoriteEntity)
}

func (f *favoriteApp) getClientByID(clientID uint) error {
	_, err := f.clientRepo.FindByID(clientID)
	return err
}

func (f *favoriteApp) getProductByID(productID uint) error {
	var responseDTO dto.ProductResponseDTO

	productKey := fmt.Sprintf("product-detail-%d", productID)

	if err := f.cacheRepo.Get(productKey, &responseDTO); err == nil {
		return nil
	}

	productEntity, err := f.productRepo.Detail(productID)
	if err != nil {
		return err
	}

	responseDTO = dto.ProductEntityToProductResponseDTO(productEntity)

	return f.cacheRepo.Set(productKey, responseDTO, productTimeExpiration*time.Minute)
}

// Delete implements FavoriteApp.
func (f *favoriteApp) Delete(request dto.FavoriteRequestDTO) error {
	return f.favoriteRepo.Delete(request.ToEntity())
}
