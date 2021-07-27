package main

import (
	"linweb"
	"linweb/examples/blog_example/controllers"
)

func main() {
	l := linweb.NewLinweb()
	l.AddControllers(&controllers.UserController{}, &controllers.BlogController{})
	l.Run(":9999")
}
