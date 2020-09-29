package main

import "time"

type Message struct {
	tableName  struct{}               `pg:"space.messages"`
	ID         int                    `json:"id" param:"id"`
	From       string                 `json:"from"` /*`param:"fromID"`*/
	To         string                 `json:"to,omitempty"`
	GroupID    int                    `json:"group_id,omitempty"`
	Message    map[string]interface{} `json:"message"`
	ContentURL string                 `json:"content_url"`
	Created    time.Time              `json:"created"`
}

type Content struct {
	Data []byte `json:"data"`
	Type string `json:"type"`
}
