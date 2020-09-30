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

	msg, err := h.svc.submitMessage(rawMSG)
	if err != nil {
		return err
	}

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

func (h *handler) queryMessages(c echo.Context) error {
	var req queryRequest
	err := c.Bind(&req)
	if err != nil {
		return err
	}
	msgs, err := h.svc.query(req)
	if err != nil {
		return err
	}

	target := "en"
	if validLanguage(req.Language) {
		target = *req.Language
	}

	translated, err := h.svc.translateMessages(msgs, target)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, translated)
}

var validLanguages = map[string]struct{}{
	"af":    {},
	"sq":    {},
	"am":    {},
	"ar":    {},
	"az":    {},
	"bn":    {},
	"bs":    {},
	"bg":    {},
	"zh":    {},
	"zh-TW": {},
	"hr":    {},
	"cs":    {},
	"da":    {},
	"fa-AF": {},
	"nl":    {},
	"en":    {},
	"et":    {},
	"fi":    {},
	"fr":    {},
	"fr-CA": {},
	"ka":    {},
	"de":    {},
	"el":    {},
	"ha":    {},
	"he":    {},
	"hi":    {},
	"hu":    {},
	"id":    {},
	"it":    {},
	"ja":    {},
	"ko":    {},
	"lv":    {},
	"ms":    {},
	"no":    {},
	"fa":    {},
	"ps":    {},
	"pl":    {},
	"pt":    {},
	"ro":    {},
	"ru":    {},
	"sr":    {},
	"sk":    {},
	"sl":    {},
	"so":    {},
	"es":    {},
	"es-MX": {},
	"sw":    {},
	"sv":    {},
	"tl":    {},
	"ta":    {},
	"th":    {},
	"tr":    {},
	"uk":    {},
	"ur":    {},
	"vi":    {},
}

func validLanguage(lang *string) bool {
	if lang == nil {
		return false
	}
	_, ok := validLanguages[*lang]
	return ok
}
