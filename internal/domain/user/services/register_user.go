package services

import (
	"capsuler/internal/domain/user"
	"capsuler/internal/domain/user/model"
	"errors"

	"github.com/google/uuid"
	"go.uber.org/fx"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUser struct {
	Repository user.Repository
}

func NewRegisterService(lc fx.Lifecycle, repository user.Repository) *RegisterUser {
	return &RegisterUser{
		Repository: repository,
	}
}

func (r *RegisterUser) Register(username, email, password string) error {
	if _, err := r.Repository.GetByEmail(email); err == nil {
		return errors.New("User already exists")
	}
	byteHashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := model.User{
		Id:             uuid.NewString(),
		Username:       username,
		Email:          email,
		HashedPassword: string(byteHashPassword),
	}
	return r.Repository.Save(&user)
}
