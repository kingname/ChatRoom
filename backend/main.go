package main

import (
	"github.com/kataras/iris"
	"github.com/kingname/ChatRoom/routes"
	"github.com/kataras/iris/middleware/logger"
    "github.com/kataras/iris/middleware/recover"
	)

func main(){
	app := iris.New()
	app.Use(recover.New())
    app.Use(logger.New())
	app.Controller("/", new(routes.ChatController))
	
	app.Run(iris.Addr(":8002"))
}