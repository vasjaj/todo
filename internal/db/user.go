package db

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Login        string `gorm:"unique"`
	PasswordHash string
}

func (d *Database) GetUser(login string) (*User, error) {
	user := &User{}
	if err := d.DB.Where("login = ?", login).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (d *Database) CreateUser(login string, password string) (*User, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return nil, err
	}

	user := &User{
		Login:        login,
		PasswordHash: string(bytes),
	}
	if err := d.DB.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
