package main

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/service/translate"
	"github.com/go-pg/pg/v9"
	"github.com/google/uuid"
)

type service interface {
	getMessage(id int) (Message, error)
	getMessages(globalID string) ([]Message, error)
	submitMessage(msg map[string]interface{}) (Message, error)
	query(req queryRequest) ([]Message, error)
	translateMessages(msgs []Message, target string) ([]Message, error)
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
	var msgs []Message
	_, err := s.db.Query(&msgs, `SELECT * FROM space.messages WHERE "from"=? OR "to"=? ORDER BY created DESC`, globalID, globalID)
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
	if !toOK && !groupOK {
		return Message{}, errors.New("missing or invalid receiver information")
	}

	// check for content
	var content *Content
	rawContent, ok := raw["content"]
	if ok {
		mapContent, ok := rawContent.(map[string]interface{})
		if !ok {
			return Message{}, errors.New("content key present, but invalid content structure provided")
		}
		dataString, ok := mapContent["data"].(string)
		if !ok {
			return Message{}, errors.New("invalid content data")
		}
		content = &Content{
			Data: []byte(dataString),
			Type: mapContent["type"].(string),
		}
		delete(raw, "content")
	}

	var contentURL string
	if content != nil {
		var err error
		contentURL, err = storeContent(content)
		if err != nil {
			return Message{}, err
		}
	}

	msg := &Message{
		Created:    time.Now().UTC(),
		ContentURL: contentURL,
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

func storeContent(c *Content) (url string, e error) {
	// The session the S3 Uploader will use
	sess := session.Must(session.NewSession())

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	f := base64.NewDecoder(base64.StdEncoding, bytes.NewReader(c.Data))

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("compassionspace-content"),
		Key:    aws.String(uuid.New().String() + "." + c.Type),
		Body:   f,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file, %v", err)
	}
	return result.Location, nil
}

func (s *realService) query(req queryRequest) ([]Message, error) {
	var msgs []Message
	q := s.db.Model(&msgs)

	if req.GroupID != nil {
		q.Where("group_id=?", req.GroupID)
	}

	if req.UserID != nil {
		q.Where(`"from"=? OR "to"=?`, req.UserID, req.UserID)
	}

	if req.MinimumMessageID != nil {
		q.Where("id>=?", req.MinimumMessageID)
	}

	if req.MaximumMessageID != nil {
		q.Where("id<=?", req.MaximumMessageID)
	}

	if req.Skip != nil {
		q.Offset(*req.Skip)
	}

	if req.Take != nil {
		q.Limit(*req.Take)
	}

	q.Order("id DESC")

	err := q.Select(&msgs)
	if err != nil {
		return nil, err
	}
	return msgs, nil
}

func (s *realService) translateMessages(msgs []Message, target string) ([]Message, error) {
	// The session the S3 Uploader will use
	sess := session.Must(session.NewSession())

	tr := translate.New(sess)

	var msgsToUpdate []Message
	for i := range msgs {
		msg := msgs[i]
		// does the translated value already exist?
		if msg.Message[target] != nil {
			continue
		}

		text, ok := msg.Message["text"].(map[string]interface{})
		if !ok {
			return nil, errors.New("could not parse text")
		}

		in, ok := text["source"].(string)
		if !ok {
			return nil, errors.New("message has no source text")
		}
		out, err := tr.Text(&translate.TextInput{
			Text:               aws.String(in),
			SourceLanguageCode: aws.String("auto"),
			TargetLanguageCode: aws.String(target),
		})
		if err != nil {
			return nil, err
		}
		msg.Message[target] = out.TranslatedText
		msgs[i] = msg
		msgsToUpdate = append(msgsToUpdate, msg)
	}

	go s.db.Update(&msgsToUpdate)

	return msgs, nil
}
