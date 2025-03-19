package dto_test

import (
	"testing"

	"github.com/andrersp/favorites/internal/domain/dto"
	"github.com/stretchr/testify/assert"
)

func TestLoginRequest(t *testing.T) {
	tests := []struct {
		name      string
		input     dto.LoginRequestDTO
		wantError bool
	}{
		{
			name: "login validate success",
			input: dto.LoginRequestDTO{
				Email: validEmail,
			},
			wantError: false,
		},
		{
			name: "login validate success",
			input: dto.LoginRequestDTO{
				Email: "mail@mail",
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
