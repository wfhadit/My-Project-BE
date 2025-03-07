package data

import (
	"errors"
	user "my-project-be/features/user"

	"gorm.io/gorm"
)

type model struct {
	connection *gorm.DB
}

func NewModel(db *gorm.DB) user.UserModel {
	return &model{connection: db}
}

func (m *model) Register(newData user.User) error {
	err := m.connection.Create(&newData).Error
	if err != nil {
		return errors.New("terjadi masalah pada database")
	}
	return nil
}

func (m *model) Login(email string) (user.User, error) {
	result := user.User{}
	if err := m.connection.Where("email = ?", email).Find(&result).Error; err != nil {
		return user.User{}, err
	}
	return result, nil
}