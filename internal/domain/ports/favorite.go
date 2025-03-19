package ports

import "github.com/andrersp/favorites/internal/domain/entity"

type FavoriteRepository interface {
	Add(favorite entity.Favorite) error
	Delete(favorite entity.Favorite) error
}
