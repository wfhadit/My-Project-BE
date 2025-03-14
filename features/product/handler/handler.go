package handler

import (
	"my-project-be/features/product"
	"my-project-be/helper"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type ProductController struct {
	service product.ProductService
}

func ProductHandler(s product.ProductService) *ProductController {
	return &ProductController{service: s}
}

func (pc *ProductController) CreateProduct(c echo.Context) error {
	_, errToken := c.Get("user").(*jwt.Token)
	if !errToken {
		return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusUnauthorized, "Token tidak terbaca", nil))
	}
	newProduct := ProductRequest{}
	errBind := c.Bind(&newProduct)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Salah input", nil))
	}
	result,errInsert := pc.service.CreateProduct(product.Product{ Nama: newProduct.Nama, Brand: newProduct.Brand, Category: newProduct.Category, Price: newProduct.Price, Amount: newProduct.Amount, Description: newProduct.Description, Image: newProduct.Image })
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Server gagal membaca input", nil))
	}
	response := ProductResponse{ ID: result.ID, Nama: result.Nama, Brand: result.Brand, Category: result.Category, Price: result.Price, Amount: result.Amount, Description: result.Description, Image: result.Image }
	return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "Product berhasil dibuat", response))
}

func (pc *ProductController) GetAllProduct(c echo.Context) error {
	result, err := pc.service.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Server gagal membaca input", nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "Berhasil membaca data", result))
}

func (pc *ProductController) GetProductById(c echo.Context) error {
	productID, err := strconv.ParseUint(c.Param("productID"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Salah input", nil))
	}
	result, err := pc.service.GetProductByID(uint(productID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Server gagal membaca input", nil))
	}
	response := ProductResponse{ ID: result.ID, Nama: result.Nama, Brand: result.Brand, Category: result.Category, Price: result.Price, Amount: result.Amount, Description: result.Description, Image: result.Image }
	return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "Product berhasil dibuat", response))
}