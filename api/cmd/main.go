package main

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var startTime = time.Now()

func main() {
	e := echo.New()
	e.Debug = true

	h, err := newHandler()
	if err != nil {
		e.Logger.Panic(err)
	}

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.BodyDump(func(c echo.Context, req []byte, resp []byte) {
		if len(req) > 1000 {
			req = req[:1000]
		}
		if len(resp) > 1000 {
			resp = resp[:1000]
		}
		c.Logger().Printj(map[string]interface{}{
			"req":  string(req),
			"resp": string(resp),
		})
	}))

	e.GET("/", healthCheck)

	user := e.Group("user/:from")
	{
		user.POST("/message", h.storeMessage)
		user.GET("/message/:id", h.retrieveContent)
		user.GET("/messages", h.retrieveMessages)
	}
	e.GET("/messages", h.queryMessages)

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = ":80"
	}
	e.Logger.Panic(e.Start(port))
}

type health struct {
	Start  time.Time
	Uptime string
}

func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, health{
		Start:  startTime,
		Uptime: time.Since(startTime).String(),
	})
}
