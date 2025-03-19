package dto_test

import (
	"reflect"
	"testing"

	"github.com/andrersp/favorites/internal/domain/dto"
	"github.com/andrersp/favorites/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

const (
	validEmail = "valideemail@mail.com"
	clientName = "teste name client"
)

var (
	product = dto.ProductResponseDTO{
		ID:    1,
		Image: "productImage",
		Price: 10.0,
	}
)

func TestClientRequest(t *testing.T) {
	tests := []struct {
		name      string
		input     dto.ClientRequestDTO
		wantError bool
	}{
		{
			name: "create client success",
			input: dto.ClientRequestDTO{
				Name:  clientName,
				Email: validEmail,
			},
			wantError: false,
		},
		{
			name: "create client error empty name",
			input: dto.ClientRequestDTO{
				Name:  "",
				Email: validEmail,
			},
			wantError: true,
		},
		{
			name: "create client error empty email",
			input: dto.ClientRequestDTO{
				Name: clientName,
			},
			wantError: true,
		},
		{
			name: "create client error invalid email",
			input: dto.ClientRequestDTO{
				Name:  clientName,
				Email: "invalidmail.com",
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.input.Validate()
			assert.Equal(t, err != nil, tt.wantError)
		})
	}
}

func TestClientDetailRequest(t *testing.T) {
	tests := []struct {
		name      string
		input     dto.ClientDetailRequestDTO
		wantError bool
	}{
		{
			name: "success",
			input: dto.ClientDetailRequestDTO{
				ID: 1,
			},
			wantError: false,
		},
		{
			name: "error",
			input: dto.ClientDetailRequestDTO{
				ID: 0,
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.input.Validate()
			assert.Equal(t, err != nil, tt.wantError)
		})
	}
}

func TestClientDetailFromEntity(t *testing.T) {
	tests := []struct {
		name     string
		entity   entity.Client
		wantoDTO dto.ClientDetailResponseDTO
	}{
		{
			name: "create client resume response",
			entity: entity.Client{
				ID:    1,
				Name:  clientName,
				Email: validEmail,
			},
			wantoDTO: dto.ClientDetailResponseDTO{
				ID:        1,
				Name:      clientName,
				Email:     validEmail,
				Favorites: make([]dto.ProductResponseDTO, 0),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dto := dto.ClientDetailFromEntity(tt.entity)
			if assert.True(t, reflect.DeepEqual(tt.wantoDTO, dto)) {
				dto.AddProduct(product)
				assert.Len(t, dto.Favorites, 1)
			}
		})
	}
}

func TestClientListFromEntity(t *testing.T) {
	tests := []struct {
		name     string
		entity   []entity.Client
		wantoDTO []dto.ClientResumeResponseDTO
	}{
		{
			name: "create client detail response",
			entity: []entity.Client{
				{
					ID:    1,
					Name:  clientName,
					Email: validEmail,
				}},
			wantoDTO: []dto.ClientResumeResponseDTO{
				{
					ID:    1,
					Name:  clientName,
					Email: validEmail,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dtos := dto.ClientListFromEntity(tt.entity)
			assert.True(t, reflect.DeepEqual(tt.wantoDTO, dtos))
		})
	}
}

func TestClientRequestTo(t *testing.T) {
	tests := []struct {
		name       string
		dto        dto.ClientRequestDTO
		wantEntity entity.Client
	}{
		{
			name: "request to entity",
			wantEntity: entity.Client{
				Name:  clientName,
				Email: validEmail,
			},
			dto: dto.ClientRequestDTO{
				Name:  clientName,
				Email: validEmail,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entity := tt.dto.ToEntity()
			assert.True(t, reflect.DeepEqual(tt.wantEntity, entity))
		})
	}
}
