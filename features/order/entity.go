package order

import "github.com/golang-jwt/jwt/v5"

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
	// GetOrder(token *jwt.Token)(Order, []OrderItem, error)
}

type OrderModel interface {
	CreateOrder(newOrder Order,  userid uint) (Order, error)
	// GetOrder(userid uint) (Order, []OrderItem, error)
}