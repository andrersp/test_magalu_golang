package client

import (
	"github.com/andrersp/favorites/internal/domain/entity"
	"github.com/andrersp/favorites/internal/domain/ports"
	"gorm.io/gorm"
)

type clientRepository struct {
	db *gorm.DB
}

func NewClientRepository(
	db *gorm.DB,
) ports.ClientRepository {
	repo := new(clientRepository)
	repo.db = db

	return repo
}

// Delete implements ports.ClientRepository.
func (c *clientRepository) Delete(clientID uint) error {
	return c.db.Delete(&Client{}, clientID).Error
}

// FindByID implements ports.ClientRepository.
func (c *clientRepository) FindByID(clientID uint) (entity.Client, error) {
	var clientModel Client
	err := c.db.Preload("Favorites").First(&clientModel, clientID).Error

	if err != nil {
		return entity.Client{}, err
	}

	return clientModel.ToEntity(), nil
}

// FindByEmail implements ports.ClientRepository.
func (c *clientRepository) FindByEmail(email string) (entity.Client, error) {
	var model Client

	err := c.db.First(&model, "email = ?", email).Error
	if err != nil {
		return entity.Client{}, err
	}

	return model.ToEntity(), nil
}

// List implements ports.ClientRepository.
func (c *clientRepository) List() ([]entity.Client, error) {
	models := make([]Client, 0)

	err := c.db.Find(&models).Error
	if err != nil {
		return nil, err
	}

	response := make([]entity.Client, len(models))

	for index := range models {
		response[index] = models[index].ToEntity()
	}

	return response, nil
}

// Save implements ports.ClientRepository.
func (c *clientRepository) Save(client entity.Client) error {
	clientModel := ToModel(client)
	return c.db.Create(clientModel).Error
}

// Update implements ports.ClientRepository.
func (c *clientRepository) Update(client entity.Client) error {
	clientModel := ToModel(client)

	return c.db.Save(&clientModel).Error
}
