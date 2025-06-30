package capsule

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Model struct {
	Id          string
	Name        string
	Description string
	DateToOpen  time.Time
	IsOpen      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Messages    []string
}

func NewModel(name, description string, dateToOpen time.Time) *Model {
	return &Model{
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

func (m *Model) AddMessage(message string) error {
	if m.IsOpen {
		return errors.New("capsule is already open")
	}
	m.Messages = append(m.Messages, message)
	m.UpdatedAt = time.Now()
	return nil
}

func (m *Model) Open() error {
	if m.IsOpen {
		return errors.New("capsule is already open")
	}
	m.IsOpen = true
	m.UpdatedAt = time.Now()
	return nil
}
