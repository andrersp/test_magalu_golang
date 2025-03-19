package dto_test

import (
	"testing"

	"github.com/andrersp/favorites/internal/domain/dto"
	"github.com/stretchr/testify/assert"
)

func TestFavoriteRequest(t *testing.T) {
	tests := []struct {
		name      string
		input     dto.FavoriteRequestDTO
		wantError bool
	}{
		{
			name: "save favorite success",
			input: dto.FavoriteRequestDTO{
				ClientID:  1,
				ProductID: 1,
			},
			wantError: false,
		},
		{
			name: "save favorites error client email",
			input: dto.FavoriteRequestDTO{
				ProductID: 1,
				ClientID:  0,
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

func TestFavoriteDTOToEntity(t *testing.T) {
	dto := dto.FavoriteRequestDTO{
		ClientID:  1,
		ProductID: 1,
	}

	entity := dto.ToEntity()

	assert.Equal(t, dto.ClientID, entity.ClientID)
	assert.Equal(t, dto.ProductID, entity.ProductID)
}
