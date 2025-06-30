package capsule

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Entity struct {
	Id          string
	Name        string
	Description string
	DateToOpen  time.Time
	IsOpen      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Messages    []string
}

func NewEntity(name, description string, dateToOpen time.Time) *Entity {
	return &Entity{
		Id:          uuid.NewString(),
		Name:        name,
		Description: description,
		DateToOpen:  dateToOpen,
		IsOpen:      false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Messages:    []string{},
	}
}

func (m *Entity) AddMessage(message string) error {
	if m.IsOpen {
		return errors.New("capsule is already open")
	}
	m.Messages = append(m.Messages, message)
	m.UpdatedAt = time.Now()
	return nil
}

func (m *Entity) Open() error {
	if m.IsOpen {
		return errors.New("capsule is already open")
	}
	m.IsOpen = true
	m.UpdatedAt = time.Now()
	return nil
}
