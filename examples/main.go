package main

import (
	"linweb"
	"linweb/examples/controllers"
)

func main() {
	r := linweb.New()
	r.AddControllers(&controllers.UserController{}, &controllers.BlogController{})
	r.Run(":9999")
}
