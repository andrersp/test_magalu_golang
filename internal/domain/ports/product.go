package ports

import "github.com/andrersp/favorites/internal/domain/entity"

type ProductRepository interface {
	List(page int) ([]entity.Product, error)
	Detail(productID uint) (entity.Product, error)
}
