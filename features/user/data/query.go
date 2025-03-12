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
	if err := m.connection.Where("email = ?", email).First(&result).Error; err != nil {
		return user.User{}, err
	}
	return result, nil
}

func (m *model) GetUserByID(id uint) (user.User, error) {
	result := user.User{}
	if err := m.connection.Where("id = ?", id).First(&result).Error; err != nil {
		return user.User{}, err
	}
	return result, nil
}

func (m *model) Update(id uint, newData user.User) (user.User, error) {
	if err := m.connection.Where("id = ?", id).Updates(&newData).Error; err != nil {
		return user.User{}, err
	}
	return newData, nil
}