package rest

import (
	"melvad/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSigntoModel(t *testing.T) {
	type usecase struct {
		name        string
		arg         *SignRequest
		expectValue model.HashData
		exErr       error
	}
	tests := []usecase{
		{
			name:        "ok",
			arg:         &SignRequest{Text: "text3432", Key: "123"},
			expectValue: model.HashData{Text: "text3432", Key: "123"},
		},
		{
			name:  "err",
			arg:   &SignRequest{Text: "test123", Key: ""},
			exErr: ErrBadParams,
		},
	}
	for _, usecase := range tests {
		user, err := usecase.arg.toModel()
		if err != nil {
			assert.Equal(t, usecase.exErr, err)

		}
		assert.Equal(t, usecase.expectValue, user)
	}
}
