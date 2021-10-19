package controllers

import (
	"linweb"
	"linweb/examples/blog/db"
	"linweb/interfaces"
	"net/http"
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
