package handler

import (
	"math"
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
	category, brand, sort, q := c.QueryParam("category"),c.QueryParam("brand"),c.QueryParam("price"),c.QueryParam("q")
	page,_ := strconv.Atoi(c.QueryParam("page"))
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * 10	
	result, total, err := pc.service.GetAllProducts(offset, category, brand, sort, q)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Server gagal membaca input", nil))
	}
	totalPages := int(math.Ceil(float64(total) / 10))
	response := []ProductResponse{}
	for _, v := range result {
		response = append(response, ProductResponse{ ID: v.ID, Nama: v.Nama, Brand: v.Brand, Category: v.Category, Price: v.Price, Amount: v.Amount, Description: v.Description, Image: v.Image })
	}
	return c.JSON(http.StatusCreated, helper.ResponseGetAllProducts(http.StatusCreated, "Product berhasil dibuat", totalPages,response))
}

func (pc *ProductController) GetProductById(c echo.Context) error {
	productID, err := strconv.Atoi(c.Param("productID"))
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