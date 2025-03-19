package ports

import "github.com/andrersp/favorites/internal/domain/entity"

type ClientRepository interface {
	Save(client entity.Client) error
	FindByEmail(email string) (entity.Client, error)
	FindByID(clientID uint) (entity.Client, error)
	Update(client entity.Client) error
	List() ([]entity.Client, error)
	Delete(clientID uint) error
}
