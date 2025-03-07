package services

import (
	"errors"
	"log"
	user "my-project-be/features/user"
	"my-project-be/features/user/handler"
	"my-project-be/helper"
	"my-project-be/middlewares"

	"github.com/go-playground/validator/v10"
)

type service struct {
	model user.UserModel
	pm helper.PasswordManager
	v *validator.Validate
}

func NewService(m user.UserModel) user.UserService {
	return &service{
		model: m,
		pm: helper.NewPasswordManager(),
		v: validator.New(),
	}
}

func (s *service) Register(newData user.User) error {
	registerValidate := handler.RegisterRequest{
		Nama:     newData.Nama,
		Email:    newData.Email,
		Password: newData.Password,
	}
	err := s.v.Struct(&registerValidate)
	if err != nil {
		return err
	}

	result, _ := s.pm.HashPassword(newData.Password)
	newData.Password = result

	err = s.model.Register(newData)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) Login(loginData user.User) (user.User, string, error) {
	loginValidate :=handler.LoginRequest{
		Email:    loginData.Email,
		Password: loginData.Password,
	}
	err := s.v.Struct(&loginValidate)
	if err != nil {
		log.Println("terjadi error", err.Error())
		return user.User{}, "", err
	}

	data, err := s.model.Login(loginValidate.Email)
	if err != nil {
		return user.User{}, "", err
	}

	err = s.pm.CheckPassword(loginValidate.Password, data.Password)
	if err != nil {
		return user.User{}, "", errors.New(helper.ServiceGeneralError)
	}

	token, err := middlewares.GenerateJWT(data.ID, data.Nama)
	if err != nil {
		return user.User{}, "", errors.New(helper.ServiceGeneralError)
	}

	return data, token, nil
}

