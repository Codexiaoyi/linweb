package main

import (
	"fmt"
	"linweb"
	"linweb/examples/blog/controllers"
	"linweb/interfaces"
	"log"
)

func main() {
	l := linweb.NewLinWeb()
	l.AddMiddlewares(PrintHelloMiddleware)
	l.AddControllers(&controllers.UserController{}, &controllers.BlogController{})
	err := l.Run(":4560")
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func PrintHelloMiddleware(c interfaces.IContext) {
	fmt.Println("hello linweb!")
	c.Next()
	fmt.Println("byebye linweb")
}
