package product

import "github.com/andrersp/favorites/internal/domain/entity"

type Product struct {
	Image       string  `json:"image"`
	Brand       string  `json:"brand"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	ID          uint    `json:"id"`
	ReviewScore float32 `json:"reviewScore"`
}

func (p *Product) ToEntity() entity.Product {
	return entity.Product{
		Image:       p.Image,
		Brand:       p.Brand,
		Title:       p.Title,
		Price:       p.Price,
		ID:          p.ID,
		ReviewScore: p.ReviewScore,
	}
}

func ProductListToEntity(products []Product) []entity.Product {
	result := make([]entity.Product, len(products))

	for index := range products {
		result[index] = products[index].ToEntity()
	}

	return result
}
