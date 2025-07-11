package capsule

import "capsuler/internal/domain/capsule/model"

type Repository interface {
	Save(capsule *model.Capsule) error
	GetById(id string) (*model.Capsule, error)
	Remove(id string) error
	Count() int
}
