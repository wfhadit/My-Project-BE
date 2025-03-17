package handler

import (
	"my-project-be/features/cart"
	"my-project-be/helper"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type CartController struct {
	service cart.CartService
}

func CartHandler(s cart.CartService) *CartController {
	return &CartController{service: s}
}

func (cc *CartController) AddCart(c echo.Context) error {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusUnsupportedMediaType, "salah cara ambil token", nil))
	}
	newCart := CartRequest{}
	errBind := c.Bind(&newCart)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusUnsupportedMediaType, "Salah input", nil))
	}
	errInsert := cc.service.AddCart(token, cart.Cart{ ProductID: newCart.ProductID, ProductNama: newCart.ProductNama, ProductImage: newCart.ProductImage, ProductPrice: newCart.ProductPrice, Quantity: newCart.Quantity }) 
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Server gagal membaca input", nil))
	}
	return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "Cart berhasil dibuat", errInsert))
}

func (cc *CartController) GetCart(c echo.Context) error {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusUnsupportedMediaType, "salah cara ambil token", nil))
	}
	result, err := cc.service.GetCart(token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Server gagal membaca input", nil))
	}
	response := []CartResponse{}
	for _, v := range result {
		response = append(response, CartResponse{ ProductID: v.ProductID, ProductNama: v.ProductNama, ProductImage: v.ProductImage, ProductPrice: v.ProductPrice, Quantity: v.Quantity })
	}
	return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "Cart berhasil dibuat", response))
	
}

func (cc *CartController) DeleteCartByID(c echo.Context) error {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusUnsupportedMediaType, "salah cara ambil token", nil))
	}
	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Salah ambil ID pada route", nil))
	}
	errDelete := cc.service.DeleteCartByID(token, uint(productID))
	if errDelete != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Server gagal membaca input", nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "Success", nil))
}


func (cc *CartController) DeleteCart(c echo.Context) error {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusUnsupportedMediaType, "salah cara ambil token", nil))
	}
	errDelete := cc.service.DeleteCart(token)
	if errDelete != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Server gagal membaca input", nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "Success", nil))
}
