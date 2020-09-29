package main

import (
	"errors"
	"log"
	"time"

	"github.com/go-pg/pg/v9"
)

type service interface {
	getMessage(id int) (Message, error)
	getMessages(globalID string) ([]Message, error)
	submitMessage(msg map[string]interface{}) (Message, error)
}

type realService struct {
	db *pg.DB
}

func (s *realService) getMessage(id int) (Message, error) {
	msg := Message{ID: id}
	err := s.db.Select(&msg)
	if err != nil {
		return Message{}, err
	}
	return msg, nil
}
func (s *realService) getMessages(globalID string) ([]Message, error) {
	s.db.AddQueryHook(DBLogger{})
	var msgs []Message
	_, err := s.db.Query(&msgs, "SELECT * FROM space.messages WHERE space.messages.from=? OR space.messages.to=? ORDER BY created DESC", globalID, globalID)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}
func (s *realService) submitMessage(raw map[string]interface{}) (Message, error) {
	from, ok := raw["from"].(string)
	if !ok {
		return Message{}, errors.New("missing or invalid sender information")
	}
	to, toOK := raw["to"].(string)
	group, groupOK := raw["group_id"].(float64)
	log.Println("toOK", toOK)
	log.Println("groupOK", groupOK)
	if !toOK && !groupOK {
		return Message{}, errors.New("missing or invalid receiver information")
	}

	msg := &Message{
		Created:    time.Now().UTC(),
		ContentURL: "not yet implemented",
		From:       from,
		To:         to,
		GroupID:    int(group),
		Message:    raw,
	}
	err := s.db.Insert(msg)
	if err != nil {
		return Message{}, err
	}

	return *msg, nil
}
