package dto_test

import (
	"reflect"
	"testing"

	"github.com/andrersp/favorites/internal/domain/dto"
	"github.com/andrersp/favorites/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestProductPageRequest(t *testing.T) {
	tests := []struct {
		name      string
		input     dto.ProductPageRequestDTO
		wantError bool
	}{
		{
			name: "validate success",
			input: dto.ProductPageRequestDTO{
				Page: 1,
			},
			wantError: false,
		},
		{
			name: "validate error",
			input: dto.ProductPageRequestDTO{
				Page: 0,
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

func TestProductDetailRequest(t *testing.T) {
	tests := []struct {
		name      string
		input     dto.ProductDetailRequestDTO
		wantError bool
	}{
		{
			name: "validate success",
			input: dto.ProductDetailRequestDTO{
				ID: 1,
			},
			wantError: false,
		},
		{
			name: "validate error",
			input: dto.ProductDetailRequestDTO{
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

func TestProductListEntityToProductResponseDTO(t *testing.T) {
	tests := []struct {
		name     string
		entity   []entity.Product
		wantoDTO []dto.ProductResponseDTO
	}{
		{
			name: "to entity",
			entity: []entity.Product{
				{
					ID:          1,
					Image:       "image",
					Price:       1.0,
					Brand:       "brand",
					Title:       "title",
					ReviewScore: 1.0,
				},
			},
			wantoDTO: []dto.ProductResponseDTO{
				{
					ID:          1,
					Image:       "image",
					Price:       1.0,
					Brand:       "brand",
					Title:       "title",
					ReviewScore: 1.0,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dtos := dto.ProductListEntityToProductResponseDTO(tt.entity)
			assert.True(t, reflect.DeepEqual(tt.wantoDTO, dtos))
		})
	}
}
