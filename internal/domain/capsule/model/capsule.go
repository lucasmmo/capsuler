package model

import (
	"errors"
	"time"
)

type Capsule struct {
	Id          string
	Name        string
	Description string
	DateToOpen  time.Time
	IsOpen      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Messages    []Message
	OwnerId     string
	PaymentId   string
}

func (c *Capsule) Open(userId string) error {
	if c.DateToOpen.Before(time.Now()) {
		return errors.New("invalid date to open the capsule")
	}
	if userId != c.OwnerId {
		return errors.New("user id is not the owner id")
	}
	c.IsOpen = true
	return nil
}

func (c *Capsule) AddMessage(message Message) error {
	if c.IsOpen {
		return errors.New("this capsule is already opened")
	}
	c.Messages = append(c.Messages, message)
	return nil
}
