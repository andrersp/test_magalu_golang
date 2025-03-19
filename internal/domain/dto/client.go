package dto

import (
	"strings"

	"github.com/andrersp/favorites/internal/domain/entity"
	"github.com/go-playground/validator/v10"
)

type ClientRequestDTO struct {
	Name  string `jaon:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

func (c *ClientRequestDTO) ToEntity() entity.Client {
	return entity.Client{
		Name:  strings.TrimSpace(c.Name),
		Email: strings.TrimSpace(c.Email),
	}
}

func (c *ClientRequestDTO) Validate() error {
	return validator.New().Struct(c)
}

type ClientDetailRequestDTO struct {
	ID uint `param:"clientID" validate:"required"`
}

func (c *ClientDetailRequestDTO) Validate() error {
	return validator.New().Struct(c)
}

type ClientResumeResponseDTO struct {
	Name  string `jaon:"name"`
	Email string `json:"email"`
	ID    uint   `json:"id"`
}

func ClientListFromEntity(clients []entity.Client) []ClientResumeResponseDTO {
	response := make([]ClientResumeResponseDTO, len(clients))

	for index := range clients {
		response[index] = ClientResumeResponseDTO{
			ID:    clients[index].ID,
			Name:  clients[index].Name,
			Email: clients[index].Email,
		}
	}

	return response
}

type ClientDetailResponseDTO struct {
	Name      string               `json:"name"`
	Email     string               `json:"email"`
	Favorites []ProductResponseDTO `json:"favorites"`
	ID        uint                 `json:"id"`
}

func (c *ClientDetailResponseDTO) AddProduct(product ProductResponseDTO) {
	c.Favorites = append(c.Favorites, product)
}

func ClientDetailFromEntity(client entity.Client) ClientDetailResponseDTO {
	return ClientDetailResponseDTO{
		ID:        client.ID,
		Name:      client.Name,
		Email:     client.Email,
		Favorites: make([]ProductResponseDTO, 0),
	}
}
