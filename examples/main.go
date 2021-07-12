package main

import (
	"linweb"
	"linweb/interfaces"
	"net/http"
)

func main() {
	r := linweb.New()
	r.GET("/", func(c interfaces.IContext) {
		c.Response().HTML(http.StatusOK, "<h1>Hello linweb</h1>")
	})
	r.GET("/hello", func(c interfaces.IContext) {
		// expect /hello?name=lingo
		c.Response().String(http.StatusOK, "hello %s, you're at %s\n", c.Request().Query("name"), c.Request().Path())
	})

	r.POST("/login", func(c interfaces.IContext) {
		c.Response().JSON(http.StatusOK, map[string]string{
			"username": c.Request().PostForm("username"),
			"password": c.Request().PostForm("password"),
		})
	})

	r.Run(":9999")
}
