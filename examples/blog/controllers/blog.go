package controllers

import (
	"linweb"
	"linweb/interfaces"
	"net/http"
	"strconv"
	"time"
)

type BlogController struct {
}

type Int int

func (d Int) Len() int {
	return 1
}

//[GET("/blog/:id")]
func (blog *BlogController) GetBlog(c interfaces.IContext) {
	id, _ := strconv.Atoi(c.Request().Param("id"))
	linweb.Cache.AddWithExpire("id", Int(id), 10*time.Second)
	c.Response().String(http.StatusOK, "id=%s", c.Request().Param("id"))
}
