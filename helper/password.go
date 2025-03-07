package helper

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type PasswordManager interface {
	HashPassword(password string) (string, error)
	CheckPassword(password string, dbPassword string) error
}

type passwordManager struct {
}

func NewPasswordManager() PasswordManager {
	return &passwordManager{}
}

func (pm *passwordManager) HashPassword(password string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("terjadi error", err.Error())
		return "", err
	}
	return string(result), nil
}

func (pm *passwordManager) CheckPassword(password string, dbPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))
	if err != nil {
		log.Println("terjadi error", err.Error())
		return err
	}
	return nil
}