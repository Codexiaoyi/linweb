package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModel(t *testing.T) {
	type RegisterDto struct {
		Name     string `validate:"min=4"`
		Password string `validate:"min=16,max=18"`
		Email    string `validate:"email"`
	}

	type UserModel struct {
		Name     string
		Age      int
		Password string
		Email    string
	}

	dto := &RegisterDto{
		Name:     "test",
		Password: "12345678912345678",
		Email:    "123456@aa.com",
	}

	user := &UserModel{}
	err := New(dto).Validate().MapToByFieldName(user).ModelError()
	assert.Empty(t, err)
	assert.Equal(t, "test", user.Name)
}
