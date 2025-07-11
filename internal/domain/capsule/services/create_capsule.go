package services

import (
	"capsuler/internal/domain/capsule"
	"capsuler/internal/domain/capsule/model"
	"errors"
	"time"

	"github.com/google/uuid"
)

type CreateCapsule struct {
	capsuleRepository capsule.Repository
}

func (c *CreateCapsule) Create(name, description, ownerId, paymentId string, dateToOpen time.Time) error {

	if dateToOpen.Before(time.Now()) {
		return errors.New("invalid date to create the capsule")
	}
	capsule := model.Capsule{
		Id:          uuid.NewString(),
		Name:        name,
		Description: description,
		DateToOpen:  dateToOpen,
		IsOpen:      false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Messages:    []model.Message{},
		OwnerId:     ownerId,
		PaymentId:   paymentId,
	}
	return c.capsuleRepository.Save(&capsule)
}
