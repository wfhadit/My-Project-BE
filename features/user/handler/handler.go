package handler

import (
	"mime/multipart"
	cart "my-project-be/features/cart/handler"
	order "my-project-be/features/order/handler"
	user "my-project-be/features/user"
	"my-project-be/helper"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
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
	errInsert := ct.service.Register(user.User{ Nama: newUser.Nama, Email: newUser.Email, Password: newUser.Password})
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
	result, token, cartResult,orderResult,err := ct.service.Login(user.User{Email: input.Email, Password: input.Password})
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusUnsupportedMediaType,helper.UserInputError, nil))
	}
	responseCart := []cart.CartResponse{}
	for _, v := range cartResult {
		responseCart = append(responseCart, cart.CartResponse{ProductID: v.ProductID, ProductNama: v.ProductNama, ProductImage: v.ProductImage, ProductPrice: v.ProductPrice, Quantity: v.Quantity})
	}
	responseOrder := []order.OrderResponse{}
	for _, v := range orderResult {
		responseOrderItems := []order.OrderItemResponse{}
		for _, w := range v.Items {
			responseOrderItems = append(responseOrderItems, order.OrderItemResponse{ID: w.ID, ProductID: w.ProductID, ProductName: w.ProductName, ProductImage: w.ProductImage, ProductPrice: w.ProductPrice, Quantity: w.Quantity})
		}
		responseOrder = append(responseOrder, order.OrderResponse{ID: v.ID, OrderUniqueID: v.OrderUniqueID, UserID: v.UserID, TotalPrice: v.TotalPrice, Status: v.Status, Items: responseOrderItems, PaymentMethod: v.PaymentMethod, VANumber: v.VANumber, CreatedAt: v.CreatedAt})
	}

	responseData := LoginResponse{ ID: result.ID, Nama: result.Nama, Email: result.Email, JenisKelamin: result.JenisKelamin, TanggalLahir: result.TanggalLahir,NomorHP: result.NomorHP, Alamat: result.Alamat, Foto: result.Foto}
	return c.JSON(http.StatusOK, helper.ResponseFormatLogin(responseData, token,responseCart, responseOrder))
}

func (ct *UserController) KeepLogin(c echo.Context) error {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusUnsupportedMediaType,"salah cara ambil token", nil))
	}
	result, newToken,cartResult,orderResult,err := ct.service.KeepLogin(token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusUnsupportedMediaType,helper.UserInputError, nil))
	}
	responseCart := []cart.CartResponse{}
	for _, v := range cartResult {
		responseCart = append(responseCart, cart.CartResponse{ ProductID: v.ProductID, ProductNama: v.ProductNama, ProductImage: v.ProductImage, ProductPrice: v.ProductPrice, Quantity: v.Quantity })
	}

	responseOrder := []order.OrderResponse{}
	for _, v := range orderResult {
		responseOrderItems := []order.OrderItemResponse{}
		for _, w := range v.Items {
			responseOrderItems = append(responseOrderItems, order.OrderItemResponse{ID: w.ID, ProductID: w.ProductID, ProductName: w.ProductName, ProductImage: w.ProductImage, ProductPrice: w.ProductPrice, Quantity: w.Quantity})
		}
		responseOrder = append(responseOrder, order.OrderResponse{ID: v.ID, OrderUniqueID: v.OrderUniqueID, UserID: v.UserID, TotalPrice: v.TotalPrice, Status: v.Status, Items: responseOrderItems , PaymentMethod: v.PaymentMethod, VANumber: v.VANumber, CreatedAt: v.CreatedAt})
		
	}
	responseData := LoginResponse{ ID: result.ID,Nama: result.Nama, Email: result.Email, JenisKelamin: result.JenisKelamin, TanggalLahir: result.TanggalLahir,NomorHP: result.NomorHP, Alamat: result.Alamat, Foto: result.Foto }
	return c.JSON(http.StatusOK, helper.ResponseFormatLogin(responseData, newToken, responseCart, responseOrder))
}

func (ct *UserController) Update(c echo.Context) error {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusUnsupportedMediaType,"salah cara ambil token", nil))
	}

	input := UpdateRequest{}
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusUnsupportedMediaType,"gagal input dalam struct", nil))
	}

	var file *multipart.FileHeader
	file, err := c.FormFile("file")
	if err == http.ErrMissingFile {
		file = nil 
	} else if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusUnsupportedMediaType, "error ambil file", nil))
	}

	result, err := ct.service.Update(token, user.User{ Nama: input.Nama, Email: input.Email, JenisKelamin: input.JenisKelamin, TanggalLahir: input.TanggalLahir, NomorHP: input.NomorHP, Alamat: input.Alamat }, file)
	if err != nil  {
		return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusUnsupportedMediaType,"error dalam service", nil))
	}
	responseData := UpdateResponse{ Nama: result.Nama, Email: result.Email, TanggalLahir: result.TanggalLahir, JenisKelamin: result.JenisKelamin, NomorHP: result.NomorHP, Alamat: result.Alamat, Foto: result.Foto}
	return c.JSON(http.StatusOK, responseData)

}