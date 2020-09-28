package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Debug = true

	e.Use(middleware.Logger())

	user := e.Group("user/:fromID")
	{
		user.POST("/message", storeMessage)
		user.GET("/message/:id", retrieveContent)
		user.GET("/messages", retrieveMessages)
	}

	e.Logger.Panic(e.Start(":8080"))
}
