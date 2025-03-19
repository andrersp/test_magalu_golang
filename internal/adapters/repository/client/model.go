package client

import (
	"github.com/andrersp/favorites/internal/adapters/repository/favorite"
	"github.com/andrersp/favorites/internal/domain/entity"
)

type Client struct {
	Favorites []favorite.Favorite `gorm:"foreignKey:ClientID"`
	Name      string
	Email     string
	ID        uint
}

func (c *Client) ToEntity() entity.Client {
	favorites := make([]entity.Favorite, len(c.Favorites))

	for index := range c.Favorites {
		favorites[index] = c.Favorites[index].ToEntity()
	}

	return entity.Client{
		ID:        c.ID,
		Name:      c.Name,
		Email:     c.Email,
		Favorites: favorites,
	}
}

func ToModel(client entity.Client) *Client {
	return &Client{
		ID:    client.ID,
		Name:  client.Name,
		Email: client.Email,
	}
}
