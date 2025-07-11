package services

import (
	"capsuler/internal/domain/user"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"go.uber.org/fx"
	"golang.org/x/crypto/bcrypt"
)

type LoginUser struct {
	UserRepository user.Repository
}

func NewLoginService(lc fx.Lifecycle, repository user.Repository) *LoginUser {
	return &LoginUser{
		UserRepository: repository,
	}
}

type Token string

func (s *LoginUser) Login(email, password string) (Token, error) {
	user, err := s.UserRepository.GetByEmail(email)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password)); err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.Id,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return Token(tokenString), nil
}
