package product

import (
	"errors"
	"slices"

	"github.com/andrersp/favorites/internal/domain/entity"
	"github.com/andrersp/favorites/internal/domain/ports"
)

const (
	pageLimit      = 2
	productPerPage = 10
)

var (
	products = []Product{
		{Image: "image45.jpg", Brand: "Brand7", Title: "Product Title 12", Price: 76.34, ID: 1, ReviewScore: 4.2},
		{Image: "image78.jpg", Brand: "Brand3", Title: "Product Title 89", Price: 45.67, ID: 2, ReviewScore: 3.8},
		{Image: "image23.jpg", Brand: "Brand1", Title: "Product Title 34", Price: 89.12, ID: 3, ReviewScore: 4.5},
		{Image: "image56.jpg", Brand: "Brand9", Title: "Product Title 67", Price: 23.45, ID: 4, ReviewScore: 4.0},
		{Image: "image12.jpg", Brand: "Brand2", Title: "Product Title 78", Price: 56.78, ID: 5, ReviewScore: 3.7},
		{Image: "image34.jpg", Brand: "Brand5", Title: "Product Title 45", Price: 67.89, ID: 6, ReviewScore: 4.1},
		{Image: "image67.jpg", Brand: "Brand8", Title: "Product Title 56", Price: 34.56, ID: 7, ReviewScore: 4.3},
		{Image: "image89.jpg", Brand: "Brand4", Title: "Product Title 23", Price: 78.90, ID: 8, ReviewScore: 3.9},
		{Image: "image90.jpg", Brand: "Brand6", Title: "Product Title 90", Price: 12.34, ID: 9, ReviewScore: 4.6},
		{Image: "image11.jpg", Brand: "Brand0", Title: "Product Title 11", Price: 90.12, ID: 10, ReviewScore: 3.5},
		{Image: "image22.jpg", Brand: "Brand7", Title: "Product Title 22", Price: 65.43, ID: 11, ReviewScore: 4.4},
		{Image: "image33.jpg", Brand: "Brand3", Title: "Product Title 33", Price: 54.32, ID: 12, ReviewScore: 3.6},
		{Image: "image44.jpg", Brand: "Brand1", Title: "Product Title 44", Price: 43.21, ID: 13, ReviewScore: 4.7},
		{Image: "image55.jpg", Brand: "Brand9", Title: "Product Title 55", Price: 32.10, ID: 14, ReviewScore: 3.4},
		{Image: "image66.jpg", Brand: "Brand2", Title: "Product Title 66", Price: 21.09, ID: 15, ReviewScore: 4.8},
		{Image: "image77.jpg", Brand: "Brand5", Title: "Product Title 77", Price: 98.76, ID: 16, ReviewScore: 3.3},
		{Image: "image88.jpg", Brand: "Brand8", Title: "Product Title 88", Price: 87.65, ID: 17, ReviewScore: 4.9},
		{Image: "image99.jpg", Brand: "Brand4", Title: "Product Title 99", Price: 76.54, ID: 18, ReviewScore: 3.2},
		{Image: "image10.jpg", Brand: "Brand6", Title: "Product Title 10", Price: 65.43, ID: 19, ReviewScore: 4.0},
		{Image: "image20.jpg", Brand: "Brand0", Title: "Product Title 20", Price: 54.32, ID: 20, ReviewScore: 3.1},
	}
)

type fakeProductRepository struct {
}

func NewFakeProductRepository() ports.ProductRepository {
	repo := new(fakeProductRepository)

	return repo
}

// Detail implements ports.ProdutoRepository.
func (f *fakeProductRepository) Detail(productID uint) (entity.Product, error) {
	if index := slices.IndexFunc(products, func(product Product) bool {
		return productID == product.ID
	}); index >= 0 {
		return products[index].ToEntity(), nil
	}

	return entity.Product{}, errors.New("product not found")
}

// List implements ports.ProdutoRepository.
func (f *fakeProductRepository) List(page int) ([]entity.Product, error) {
	result := make([]entity.Product, productPerPage)
	startIndex := 0
	endIndex := 10

	if page >= pageLimit {
		startIndex = 10
		endIndex = 20
	}

	x := 0

	for startIndex := startIndex; startIndex < endIndex; startIndex++ {
		result[x] = products[startIndex].ToEntity()
		x++
	}

	return result, nil
}
