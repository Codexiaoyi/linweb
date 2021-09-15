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

//[GET("/blog/:id")]
func (blog *BlogController) GetBlog(c interfaces.IContext) {
	id, _ := strconv.Atoi(c.Request().Param("id"))
	linweb.Cache.AddWithExpire("id", id, 10*time.Second)
	c.Response().String(http.StatusOK, "id=%s", c.Request().Param("id"))
}
