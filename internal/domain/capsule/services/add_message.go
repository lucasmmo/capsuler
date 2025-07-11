package services

import (
	"capsuler/internal/domain/capsule"
	"capsuler/internal/domain/capsule/model"
	"time"

	"github.com/google/uuid"
)

type AddMessage struct {
	repository capsule.Repository
}

func (s *AddMessage) Add(caspuleId, content, userId string) (*model.Capsule, error) {
	capsule, err := s.repository.GetById(caspuleId)
	if err != nil {
		return nil, err
	}
	message := model.Message{
		Id:        uuid.NewString(),
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserId:    userId,
	}
	if err := capsule.AddMessage(message); err != nil {
		return nil, err
	}
	if err := s.repository.Save(capsule); err != nil {
		return nil, err
	}
	return capsule, nil
}
