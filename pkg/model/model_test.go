//Copyright 2021 Codexiaoyi
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

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
	customizeModel := &Model{}
	user := &UserModel{}
	err := customizeModel.New(dto).Validate().MapToByFieldName(user).ModelError()
	assert.Empty(t, err)
	assert.Equal(t, "test", user.Name)
}
