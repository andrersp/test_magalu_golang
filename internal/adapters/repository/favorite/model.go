package favorite

import "github.com/andrersp/favorites/internal/domain/entity"

type Favorite struct {
	ID        uint
	ClientID  uint
	ProductID uint
}

func (f *Favorite) ToEntity() entity.Favorite {
	return entity.Favorite{
		ClientID:  f.ClientID,
		ProductID: f.ProductID,
	}
}

func ToModel(favorite entity.Favorite) Favorite {
	return Favorite{
		ClientID:  favorite.ClientID,
		ProductID: favorite.ProductID,
	}
}
