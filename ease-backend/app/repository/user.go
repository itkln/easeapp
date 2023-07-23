package repository

import (
	"github.com/itkln/ease-backend/app/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) GetUser(user *models.UserAuth) {

}
