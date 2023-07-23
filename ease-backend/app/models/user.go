package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID        uint
	Name      string
	Surname   string
	Email     string
	Password  string
	CreatedAt time.Time
	DeletedAt time.Time
}

type UserAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
