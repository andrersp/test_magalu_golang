package app

import (
	"github.com/andrersp/favorites/internal/domain/dto"
	"github.com/andrersp/favorites/internal/domain/entity"
	"github.com/andrersp/favorites/internal/domain/ports"
)

const (
	adminEmail = "admin@admin.com"
)

type LoginApp interface {
	Login(email string) (dto.LoginResponseDTO, error)
}

type loginApp struct {
	clientReposo    ports.ClientRepository
	securityService ports.Security
}

func NewLoginApp(clientReposo ports.ClientRepository, securityService ports.Security) LoginApp {
	return &loginApp{
		clientReposo:    clientReposo,
		securityService: securityService,
	}
}

func (l *loginApp) Login(email string) (dto.LoginResponseDTO, error) {
	if email == adminEmail {
		return l.adminLogin()
	}

	_, err := l.clientReposo.FindByEmail(email)
	if err != nil {
		return dto.LoginResponseDTO{}, err
	}

	token, err := l.securityService.CreateToken(entity.CLIENT)
	if err != nil {
		return dto.LoginResponseDTO{}, err
	}

	return dto.LoginResponseDTO{Token: token}, nil
}

func (l *loginApp) adminLogin() (dto.LoginResponseDTO, error) {
	token, err := l.securityService.CreateToken(entity.ADMIN)
	if err != nil {
		return dto.LoginResponseDTO{}, err
	}

	return dto.LoginResponseDTO{Token: token}, nil
}
