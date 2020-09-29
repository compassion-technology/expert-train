package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func newHandler() (*handler, error) {
	conn, ok := os.LookupEnv("CONN")
	if !ok {
		return nil, errors.New("missing configuration parameter")
	}
	db, err := setup(conn)
	if err != nil {
		return nil, fmt.Errorf("could not setup db: %s", err.Error())
	}

	return &handler{
		svc: &realService{
			db: db,
		},
	}, nil
}

type handler struct {
	svc service
}

func (h *handler) storeMessage(c echo.Context) error {
	rawMSG := make(map[string]interface{})
	err := c.Bind(&rawMSG)
	if err != nil {
		return err
	}

	// rawMSG["From"] = c.Param("fromID")
	c.Logger().Debugj(map[string]interface{}{
		"rawMSG": rawMSG,
	})
	msg, err := h.svc.submitMessage(rawMSG)
	if err != nil {
		return err
	}

	c.Logger().Debugj(map[string]interface{}{
		"messages": messages,
	})
	return c.JSON(http.StatusCreated, msg)
}

func (h *handler) retrieveContent(c echo.Context) error {
	var req Message
	err := c.Bind(&req)
	if err != nil {
		return err
	}
	msg, err := h.svc.getMessage(req.ID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, msg)
}

func (h *handler) retrieveMessages(c echo.Context) error {
	var req Message
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	msgs, err := h.svc.getMessages(req.From)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, msgs)
}
