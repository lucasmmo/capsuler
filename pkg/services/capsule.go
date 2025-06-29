package services

import (
	"capsuler/pkg/capsule"
	"fmt"
	"time"
)

type CapsuleRepository interface {
	Save(*capsule.Capsule) error
	GetById(id string) (*capsule.Capsule, error)
}

func CreateCapsule(name, description string, dateToOpen time.Time, capsuleRepository CapsuleRepository) (string, error) {
	newCapsule := capsule.Builder().
		WithName(name).
		WithDescription(description).
		WithDateToOpen(dateToOpen).
		Build()

	if err := capsuleRepository.Save(newCapsule); err != nil {
		return "", fmt.Errorf("error when create capsule in service: %v", err)
	}

	return newCapsule.GetId(), nil
}

func OpenCapsule(id string, capsuleRepository CapsuleRepository) error {
	capsule, err := capsuleRepository.GetById(id)
	if err != nil {
		return fmt.Errorf("error when getting capsule in service: %v", err)
	}

	if err := capsule.Open(); err != nil {
		return fmt.Errorf("error when open the capsule in service: %v", err)
	}

	if err := capsuleRepository.Save(capsule); err != nil {
		return fmt.Errorf("error when create capsule in service: %v", err)
	}
	return nil
}

func AddMessageToCapsule(id, message string, capsuleRepository CapsuleRepository) (*capsule.Capsule, error) {
	capsule, err := capsuleRepository.GetById(id)
	if err != nil {
		return nil, fmt.Errorf("error when getting capsule in service: %v", err)
	}

	if err := capsule.AddMessage(message); err != nil {
		return nil, fmt.Errorf("error when added message in service: %v", err)
	}

	if err := capsuleRepository.Save(capsule); err != nil {
		return capsule, fmt.Errorf("error when create capsule in service: %v", err)
	}
	return capsule, nil

}
