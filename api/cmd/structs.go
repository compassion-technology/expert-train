package main

import "time"

type Message struct {
	ID         int    `param:"id"`
	From       string `param:"fromID"`
	To         string
	GroupID    string
	Text       string
	ContentURL string
	Created    time.Time
}
