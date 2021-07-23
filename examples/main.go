package main

import (
	"linweb"
	"linweb/examples/controllers"
)

func main() {
	r := linweb.NewLinweb()
	r.AddControllers(&controllers.UserController{}, &controllers.BlogController{})
	r.Run(":9999")
}
