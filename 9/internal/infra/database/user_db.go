package database

import (
	"api/internal/entity"

	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{DB: db}
}

func (u *User) Create(user *entity.User) error {
	return u.DB.Create(user).Error
}

func (u *User) FindByEmail(emailId string) (*entity.User, error) {
	var userFound entity.User
	err := u.DB.First(&userFound, "email = ?", emailId).Error
	return &userFound, err
}
