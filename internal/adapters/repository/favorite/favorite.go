package favorite

import (
	"github.com/andrersp/favorites/internal/domain/entity"
	"github.com/andrersp/favorites/internal/domain/ports"
	"gorm.io/gorm"
)

type favoriteRepository struct {
	db *gorm.DB
}

func NewFavoriteRepository(
	db *gorm.DB,
) ports.FavoriteRepository {
	repo := new(favoriteRepository)
	repo.db = db

	return repo
}

// Add implements ports.FavoriteRepository.
func (f *favoriteRepository) Add(favorite entity.Favorite) error {
	model := ToModel(favorite)
	return f.db.Create(&model).Error
}

// Delete implements ports.FavoriteRepository.
func (f *favoriteRepository) Delete(favorite entity.Favorite) error {
	favoriteModel := ToModel(favorite)
	return f.db.Where("client_id = ? AND product_id = ?",
		favoriteModel.ClientID, favoriteModel.ProductID).Delete(&Favorite{}).Error
}
