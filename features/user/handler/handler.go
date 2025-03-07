package handler

import (
	user "my-project-be/features/user"
	"my-project-be/helper"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

const (
	bindError = "error bind"
)

type UserController struct {
	service user.UserService
}

func NewUserHandler(s user.UserService) *UserController {
	return &UserController{service: s}
}

func (ct *UserController) Register(c echo.Context) error {
	newUser := RegisterRequest{}
	errBind := c.Bind(&newUser)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusUnsupportedMediaType,bindError+errBind.Error(), nil))
	}

	inputUser := user.User{
		Nama:     newUser.Nama,
		Email:    newUser.Email,
		Password: newUser.Password,
	}
	errInsert := ct.service.Register(inputUser)
	if errInsert != nil {
		if strings.Contains(errInsert.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusUnsupportedMediaType,helper.UserInputError, nil))
		}
		return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusUnsupportedMediaType,helper.UserInputError, nil))
	}
	return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "Register berhasil, silahkan login",nil))
}

func (ct *UserController) Login(c echo.Context) error {
	input := LoginRequest{}
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusUnsupportedMediaType,bindError+errBind.Error(), nil))
	}
	result, token, err := ct.service.Login(user.User{Email: input.Email, Password: input.Password})
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusUnsupportedMediaType,helper.UserInputError, nil))
	}
	responseData := LoginResponse{Nama: result.Nama, Email: result.Email, Token: token}
	return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "Login berhasil", responseData))
}