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

type RegisterDto struct {
	Name     string `json:"Name"`
	Password string `json:"Password"`
	Email    string `json:"email"`
}

type UserController struct {
}

//[POST("/login")]
func (user *UserController) Login(c interfaces.IContext, dto LoginDto) {
	fmt.Println(dto)
	c.Response().HTML(http.StatusOK, "<h1>Hello linweb</h1>")
}

//[POST("/register")]
func (user *UserController) Register(c interfaces.IContext, dto RegisterDto) {
	fmt.Println(dto)
	c.Response().HTML(http.StatusOK, "<h1>Hello linweb</h1>")
}
