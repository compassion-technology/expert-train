package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func storeMessage(c echo.Context) error {
	var msg Message
	err := c.Bind(&msg)
	if err != nil {
		return err
	}

	err = putMessage(&msg)
	if err != nil {
		return err
	}

	c.Logger().Debugj(map[string]interface{}{
		"messages": messages,
	})
	return c.JSON(http.StatusCreated, msg)
}

func retrieveContent(c echo.Context) error {
	var req Message
	err := c.Bind(&req)
	if err != nil {
		return err
	}
	msg, err := getMessage(req.ID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, msg)
}

func retrieveMessages(c echo.Context) error {
	var req Message
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	msgs, err := getMessages(req.From)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, msgs)
}
