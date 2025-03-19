package ports

import "github.com/andrersp/favorites/internal/domain/entity"

type Security interface {
	CreateToken(role entity.Role) (string, error)
	ValidateToken(token string) (entity.SecurityData, error)
}
