package entity

import (
	"api/pkg/entity"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        entity.ID
	Name      string
	Email     string
	Passsword string
}

func NewUser(name, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:        entity.NewID(),
		Name:      name,
		Email:     email,
		Passsword: string(hash),
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Passsword), []byte(password))
	return err == nil
}
