package app

import (
	"fmt"
	"time"

	"github.com/andrersp/favorites/internal/domain/dto"
	"github.com/andrersp/favorites/internal/domain/ports"
)

const (
	productTimeExpiration = 20
)

type ProductApp interface {
	Detail(productID uint) (dto.ProductResponseDTO, error)
	List(page int) ([]dto.ProductResponseDTO, error)
}

type productApp struct {
	productRepository ports.ProductRepository
	cacheRepository   ports.CacheRepository
}

// Detail implements ProductApp.
func (p *productApp) Detail(productID uint) (dto.ProductResponseDTO, error) {
	var responseDTO dto.ProductResponseDTO

	productKey := fmt.Sprintf("product-detail-%d", productID)

	if err := p.cacheRepository.Get(productKey, &responseDTO); err == nil {
		return responseDTO, nil
	}

	response, err := p.productRepository.Detail(productID)
	if err != nil {
		return dto.ProductResponseDTO{}, err
	}

	responseDTO = dto.ProductEntityToProductResponseDTO(response)

	_ = p.cacheRepository.Set(productKey, responseDTO, productTimeExpiration*time.Minute)

	return dto.ProductEntityToProductResponseDTO(response), nil
}

// List implements ProductApp.
func (p *productApp) List(page int) ([]dto.ProductResponseDTO, error) {
	var response []dto.ProductResponseDTO

	productKey := fmt.Sprintf("product-page-%d", page)

	if err := p.cacheRepository.Get(productKey, &response); err == nil {
		return response, nil
	}

	products, err := p.productRepository.List(page)
	if err != nil {
		return nil, err
	}

	response = dto.ProductListEntityToProductResponseDTO(products)

	_ = p.cacheRepository.Set(productKey, response, productTimeExpiration*time.Minute)

	return response, nil
}

func NewProductApp(
	productRepository ports.ProductRepository,
	cacheRepository ports.CacheRepository,
) ProductApp {
	app := new(productApp)
	app.productRepository = productRepository
	app.cacheRepository = cacheRepository

	return app
}
