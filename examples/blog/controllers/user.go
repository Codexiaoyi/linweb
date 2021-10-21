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

package controllers

import (
	"net/http"

	"github.com/Codexiaoyi/linweb"
	"github.com/Codexiaoyi/linweb/examples/blog/db"
	"github.com/Codexiaoyi/linweb/interfaces"
)

type RegisterDto struct {
	Name     string `json:"Name"`
	Password string `json:"Password"`
	Address  string `json:"Address"`
}

type LoginDto struct {
	Name     string `json:"Name"`
	Password string `json:"Password"`
}

type UserController struct {
	UserRepo *db.UserRepository
}

//[POST("/register")]
func (user *UserController) Register(c interfaces.IContext, dto RegisterDto) {
	dataModel := &db.UserModel{Name: "aaa", Password: "aaa"}
	err := linweb.NewModel(&dto).Validate().MapToByFieldName(dataModel).ModelError()
	if err != nil {
		c.Response().String(http.StatusInternalServerError, "Model error :%s!", err.Error())
		return
	}
	user.UserRepo.AddUser(dataModel)
	c.Response().String(http.StatusOK, "Welcome %s!", dataModel.Name)
}

//[POST("/login")]
func (user *UserController) Login(c interfaces.IContext, dto LoginDto) {
	dataUser := user.UserRepo.GetUserByName(dto.Name)
	if dataUser != nil && dataUser.Password == dto.Password {
		c.Response().String(http.StatusOK, "Welcome %s!", dto.Name)
		return
	}
	c.Response().String(http.StatusBadRequest, "Password error!")
}
