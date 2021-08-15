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
	dataModel := &DatabaseModel{}
	err := linweb.NewModel(dto).Validate().MapToByFieldName(dataModel).ModelError()
	if err != nil {
		c.Response().String(http.StatusInternalServerError, "Model error :%s!", err.Error())
	}
	c.Response().String(http.StatusOK, "Welcome %s!", dto.Name)
}
