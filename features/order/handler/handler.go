package handler

import (
	"my-project-be/features/order"
	"my-project-be/helper"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type OrderController struct {
	service order.OrderService
}

func OrderHandler(s order.OrderService) *OrderController {
	return &OrderController{service: s}
}

func (oc *OrderController) CreateOrder(c echo.Context) error {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusUnsupportedMediaType, "salah cara ambil token", nil))
	}
	newOrder:= OrderRequest{}
	errBind := c.Bind(&newOrder)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusUnsupportedMediaType, "Salah input", nil))
	}
	newOrderItems := []order.OrderItem{}
	for _, item := range newOrder.Items {
		newOrderItems = append(newOrderItems, order.OrderItem{ProductName: item.ProductName, ProductImage: item.ProductImage, ProductPrice: item.ProductPrice, ProductID: item.ProductID, Quantity: item.Quantity})
	}
	resultOrder, err := oc.service.CreateOrder(order.Order{TotalPrice: newOrder.TotalPrice, PaymentMethod: newOrder.PaymentMethod, Items: newOrderItems}, token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Server gagal membaca input", nil))
	}
	responseOrderItems := []OrderItemResponse{}
	for _, item := range resultOrder.Items {
		responseOrderItems = append(responseOrderItems, OrderItemResponse{ ID: item.ID, ProductID: item.ProductID, ProductName: item.ProductName, ProductImage: item.ProductImage, ProductPrice: item.ProductPrice, Quantity: item.Quantity})
	}
	responseOrder := OrderResponse{ID: resultOrder.ID, UserID: resultOrder.UserID, TotalPrice: resultOrder.TotalPrice, Status: resultOrder.Status, PaymentMethod: resultOrder.PaymentMethod, VANumber: resultOrder.VANumber, OrderUniqueID: resultOrder.OrderUniqueID, CreatedAt: resultOrder.CreatedAt,Items: responseOrderItems}
	return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "Order berhasil dibuat", responseOrder))
}