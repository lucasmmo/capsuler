package capsule

import (
	"fmt"
	"time"
)

type Service interface {
	CreateCapsule(name, description string, dateToOpen time.Time) (string, error)
	OpenCapsule(id string) error
	AddMessageToCapsule(id, message string) (*Model, error)
}

type Repository interface {
	Save(*Model) error
	GetById(id string) (*Model, error)
}

type service struct {
	capsuleRepository Repository
}

func NewService(capsuleRepository Repository) *service {
	return &service{
		capsuleRepository: capsuleRepository,
	}
}

func (s *service) CreateCapsule(name, description string, dateToOpen time.Time) (string, error) {
	capsule := NewModel(name, description, dateToOpen)

	if err := s.capsuleRepository.Save(capsule); err != nil {
		return "", fmt.Errorf("error when create capsule in service: %v", err)
	}

	return capsule.Id, nil
}

func (s *service) OpenCapsule(id string) error {
	capsule, err := s.capsuleRepository.GetById(id)
	if err != nil {
		return fmt.Errorf("error when getting capsule in service: %v", err)
	}

	if err := capsule.Open(); err != nil {
		return fmt.Errorf("error when open the capsule in service: %v", err)
	}

	if err := s.capsuleRepository.Save(capsule); err != nil {
		return fmt.Errorf("error when create capsule in service: %v", err)
	}
	return nil
}

func (s *service) AddMessageToCapsule(id, message string) (*Model, error) {
	capsule, err := s.capsuleRepository.GetById(id)
	if err != nil {
		return nil, fmt.Errorf("error when getting capsule in service: %v", err)
	}

	if err := capsule.AddMessage(message); err != nil {
		return nil, fmt.Errorf("error when added message in service: %v", err)
	}

	if err := s.capsuleRepository.Save(capsule); err != nil {
		return capsule, fmt.Errorf("error when create capsule in service: %v", err)
	}
	return capsule, nil

}
