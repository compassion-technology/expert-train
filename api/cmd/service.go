package main

import (
	"time"

	"github.com/go-pg/pg/v9"
)

type service interface {
	getMessage(id int) (Message, error)
	getMessages(globalID string) ([]Message, error)
	submitMessage(msg *Message) error
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
	_, err := s.db.Query(&msgs, "SELECT * FROM space.messages WHERE space.messages.from=? OR space.messages.to=?", globalID, globalID)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}
func (s *realService) submitMessage(msg *Message) error {
	msg.Created = time.Now().UTC()
	return s.db.Insert(msg)
}
