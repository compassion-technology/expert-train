package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Debug = true

	h, err := newHandler()
	if err != nil {
		e.Logger.Panic(err)
	}

	e.Use(middleware.Logger())

	user := e.Group("user/:from")
	{
		user.POST("/message", h.storeMessage)
		user.GET("/message/:id", h.retrieveContent)
		user.GET("/messages", h.retrieveMessages)
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = ":80"
	}
	e.Logger.Panic(e.Start(port))
}
