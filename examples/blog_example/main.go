package main

import (
	"fmt"
	"linweb"
	"linweb/examples/blog_example/controllers"
	"linweb/interfaces"
)

func main() {
	l := linweb.NewLinweb()
	l.AddMiddlewares(PrintHelloMiddleware)
	l.AddControllers(&controllers.UserController{}, &controllers.BlogController{})
	l.Run(":9999")
}

func PrintHelloMiddleware(c interfaces.IContext) {
	fmt.Println("hello linweb!")
	c.Next()
	fmt.Println("byebye linweb")
}
