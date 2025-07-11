package user

import "capsuler/internal/domain/user/model"

type Repository interface {
	Save(user *model.User) error
	GetById(id string) (*model.User, error)
	Remove(id string) error
	Count() int
	GetByEmail(email string) (*model.User, error)
}
