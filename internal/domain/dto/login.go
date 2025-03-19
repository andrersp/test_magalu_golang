package dto

import "github.com/go-playground/validator/v10"

type LoginRequestDTO struct {
	Email string `json:"email" validate:"required,email"`
}

func (l *LoginRequestDTO) Validate() error {
	return validator.New().Struct(l)
}

type LoginResponseDTO struct {
	Token string `json:"token"`
}
