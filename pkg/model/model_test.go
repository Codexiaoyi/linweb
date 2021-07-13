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
	r := &RegisterDto{
		Name:     "test",
		Password: "12345678912345678",
		Email:    "123456@aa.com",
	}
	u := &UserModel{}
	err := NewModel(r).Validate().MapTo(u).ModelError()
	assert.Empty(t, err)
	assert.Equal(t, "test", u.Name)
}
