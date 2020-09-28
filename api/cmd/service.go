package main

import (
	"errors"
	"time"
)

var messages = map[int]Message{
	1: Message{
		ID:   1,
		To:   "01234567",
		From: "09876543",
		Text: "This is a test message.",
	},
}

func getMessage(id int) (Message, error) {
	msg, ok := messages[id]
	if !ok {
		return Message{}, errors.New("not found")
	}
	return msg, nil
}

func getMessages(globalID string) ([]Message, error) {
	msgs := make([]Message, 0)
	for _, v := range messages {
		if v.To == globalID || v.From == globalID {
			msgs = append(msgs, v)
		}
	}
	return msgs, nil
}

func putMessage(msg *Message) error {
	if msg == nil {
		return errors.New("stop sending me garbage")
	}
	nextID := len(messages) + 1
	msg.ID = nextID
	msg.Created = time.Now().UTC()
	messages[nextID] = *msg
	return nil
}
