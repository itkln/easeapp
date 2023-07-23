package controllers

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/itkln/ease-backend/app/models"
	"github.com/itkln/ease-backend/app/repository"
	"time"
)

const (
	jwtSecret = "den"
)

type AuthController interface {
	GenerateJWT(user *models.UserAuth) (string, error)
	LoginUser(user *models.UserAuth) error
}

type AuthControllerImpl struct {
	repo *repository.UserRepository
}

func NewAuthServiceImpl(repo *repository.UserRepository) *AuthControllerImpl {
	return &AuthControllerImpl{repo: repo}
}

func (repo *AuthControllerImpl) GenerateJWT(user *models.UserAuth) (string, error) {
	claims := jwt.MapClaims{
		"email":     user.Email,
		"exp":       time.Now().Add(12 * time.Hour).Unix(),
		"createdAt": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

func (repo *AuthControllerImpl) LoginUser(user *models.UserAuth) error {
	return nil
}
