package services

import (
	"capsuler/internal/domain/capsule/model"

	"gorm.io/gorm"
)

type ListCapsuleByUserId struct {
	DB *gorm.DB
}

func (s *ListCapsuleByUserId) List(userId string) ([]*model.Capsule, error) {
	var capsules []*model.Capsule
	if err := s.DB.Find(&capsules, "owner_id = ?", userId).Error; err != nil {
		return nil, err
	}
	return capsules, nil
}
