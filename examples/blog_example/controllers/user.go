package controllers

import (
	"fmt"
	"linweb/interfaces"
	"net/http"
)

type LoginDto struct {
	Name     string `json:"Name"`
	Password string `json:"Password"`
}

type UserController struct {
}

//[GET("/hello")]
func (user *UserController) Hello(c interfaces.IContext) {
	c.Response().HTML(http.StatusOK, "<h1>Hello linweb</h1>")
}

//[POST("/login")]
func (user *UserController) Login(c interfaces.IContext, dto LoginDto) {
	fmt.Println(dto)
	c.Response().String(http.StatusOK, "Welcome %s!", dto.Name)
}
