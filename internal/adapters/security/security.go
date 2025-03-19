package security

import (
	"errors"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/andrersp/favorites/internal/domain/entity"
	"github.com/andrersp/favorites/internal/domain/ports"
)

const (
	roleKey             = "role"
	invalidTokenMessage = "invalid token"
)

type tokenService struct {
	expirationMinute time.Duration
	secret           string
}

// CreateToken implements ports.Security.
func (t *tokenService) CreateToken(role entity.Role) (string, error) {
	token := paseto.NewToken()

	now := time.Now()
	token.SetIssuedAt(now)
	token.SetNotBefore(now)
	token.SetExpiration(now.Add(t.expirationMinute))

	if err := token.Set(roleKey, role); err != nil {
		return "", err
	}

	key, err := t.getKey()
	if err != nil {
		return "", err
	}

	encripted := token.V4Encrypt(key, nil)

	return encripted, nil
}

// ValidateToken implements ports.Security.
func (t *tokenService) ValidateToken(token string) (entity.SecurityData, error) {
	var securityData entity.SecurityData

	key, err := t.getKey()
	if err != nil {
		return entity.SecurityData{}, err
	}

	parser := paseto.NewParser()
	tokenPayload, err := parser.ParseV4Local(key, token, nil)

	if err != nil {
		return securityData, errors.New(invalidTokenMessage)
	}

	err = tokenPayload.Get(roleKey, &securityData.Role)
	if err != nil {
		return securityData, errors.New(invalidTokenMessage)
	}

	return securityData, nil
}

func (t *tokenService) getKey() (paseto.V4SymmetricKey, error) {
	return paseto.V4SymmetricKeyFromBytes([]byte(t.secret))
}

func NewTokenService(secret string, expiration int) ports.Security {
	service := new(tokenService)
	service.secret = secret
	service.expirationMinute = time.Minute * time.Duration(expiration)

	return service
}
