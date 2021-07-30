package controllers

import (
	"linweb"
	"linweb/interfaces"
	"net/http"
)

type LoginDto struct {
	Name     string `json:"Name"`
	Password string `json:"Password"`
}

type DatabaseModel struct {
	Name     string
	Password string
}

type UserController struct {
}

//[POST("/login")]
func (user *UserController) Login(c interfaces.IContext, dto LoginDto) {
	model := linweb.NewModel(dto)
	model.Validate()
	dataModel := &DatabaseModel{}
	model.MapToByFieldName(dataModel)
	c.Response().String(http.StatusOK, "Welcome %s!", dto.Name)
}
