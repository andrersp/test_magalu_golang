package dto

import (
	"github.com/andrersp/favorites/internal/domain/entity"
	"github.com/go-playground/validator/v10"
)

type FavoriteRequestDTO struct {
	ClientID  uint `json:"clientId" param:"clientId" validate:"required"`
	ProductID uint `json:"productId" param:"productId"  validate:"required"`
}

func (f *FavoriteRequestDTO) Validate() error {
	return validator.New().Struct(f)
}

func (f *FavoriteRequestDTO) ToEntity() entity.Favorite {
	return entity.Favorite{
		ClientID:  f.ClientID,
		ProductID: f.ProductID,
	}
}
