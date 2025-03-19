package dto

import (
	"github.com/andrersp/favorites/internal/domain/entity"
	"github.com/go-playground/validator/v10"
)

type ProductPageRequestDTO struct {
	Page int `query:"page" validate:"required"`
}

func (p *ProductPageRequestDTO) Validate() error {
	return validator.New().Struct(p)
}

type ProductDetailRequestDTO struct {
	ID uint `param:"productID" validate:"required"`
}

func (p *ProductDetailRequestDTO) Validate() error {
	return validator.New().Struct(p)
}

type ProductResponseDTO struct {
	Image       string  `json:"image"`
	Brand       string  `json:"brand"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	ID          uint    `json:"id"`
	ReviewScore float32 `json:"reviewScore"`
}

func ProductEntityToProductResponseDTO(product entity.Product) ProductResponseDTO {
	return ProductResponseDTO{
		Image:       product.Image,
		Brand:       product.Brand,
		Title:       product.Title,
		Price:       product.Price,
		ID:          product.ID,
		ReviewScore: product.ReviewScore,
	}
}

func ProductListEntityToProductResponseDTO(products []entity.Product) []ProductResponseDTO {
	result := make([]ProductResponseDTO, len(products))

	for index := range products {
		result[index] = ProductEntityToProductResponseDTO(products[index])
	}

	return result
}
