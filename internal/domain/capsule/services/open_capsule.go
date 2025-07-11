package services

import (
	"capsuler/internal/domain/capsule"
)

type OpenCapsule struct {
	repository capsule.Repository
}

func (s *OpenCapsule) Open(userId, capsuleId string) error {
	capsule, err := s.repository.GetById(capsuleId)
	if err != nil {
		return err
	}
	if err := capsule.Open(userId); err != nil {
		return err
	}
	if err := s.repository.Save(capsule); err != nil {
		return err
	}
	return nil
}
