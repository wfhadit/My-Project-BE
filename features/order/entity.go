package order

import (
	"github.com/golang-jwt/jwt/v5"
)

type Order struct {
	ID            uint
	UserID        uint
	TotalPrice    uint
	Status        string
	PaymentMethod string
	VANumber      string
	OrderUniqueID string
	CreatedAt     string
	Items         []OrderItem
}

type OrderItem struct {
	ID           uint
	OrderID      uint
	ProductID    uint
	ProductName  string
	ProductImage string
	ProductPrice uint
	Quantity     uint
}

type OrderService interface {
	CreateOrder(newOrder Order, token *jwt.Token) (Order, error)
	GetOrderByUniqueID(token *jwt.Token)(Order, error)
	GetAllOrders(token *jwt.Token) ([]Order, error)
}

type OrderModel interface {
	CreateOrder(newOrder Order) (Order, error)
	GetOrderByUniqueID(uniqueID string, userid uint, newStatus string) (Order, error)
	GetAllOrders(userid uint) ([]Order, error)
	GetLastOrder(userid uint) (Order, error)
}