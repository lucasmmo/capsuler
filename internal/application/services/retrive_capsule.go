package services

import (
	"capsuler/internal/domain/capsule"
	"capsuler/internal/domain/capsule/model"
)

type RetriveCapsule struct {
	repository capsule.Repository
}

func (s *RetriveCapsule) Retrive(capsuleId string) (*model.Capsule, error) {
	return s.repository.GetById(capsuleId)
}
