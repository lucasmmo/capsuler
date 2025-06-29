package capsule

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Capsule struct {
	id          string
	name        string
	description string
	dateToOpen  time.Time
	isOpen      bool
	createdAt   time.Time
	updatedAt   time.Time
	messages    []string
}

func newCapsule() *Capsule {
	return &Capsule{
		id:        uuid.NewString(),
		isOpen:    false,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
}

func (c *Capsule) GetId() string {
	return c.id
}

func (c *Capsule) GetName() string {
	return c.name
}

func (c *Capsule) GetDescription() string {
	return c.description
}

func (c *Capsule) GetDateToOpen() time.Time {
	return c.dateToOpen
}

func (c *Capsule) WasOpened() bool {
	return c.isOpen
}

func (c *Capsule) GetMessages() []string {
	return c.messages
}

func (c *Capsule) GetCreatedAt() time.Time {
	return c.createdAt
}

func (c *Capsule) GetUpdatedAt() time.Time {
	return c.updatedAt
}

func (c *Capsule) AddMessage(message string) error {
	if c.isOpen {
		return errors.New("capsule is already open")
	}
	c.messages = append(c.messages, message)
	c.updatedAt = time.Now()
	return nil
}

func (c *Capsule) Open() error {
	if c.isOpen {
		return errors.New("capsule is already open")
	}
	c.isOpen = true
	c.updatedAt = time.Now()
	return nil
}
