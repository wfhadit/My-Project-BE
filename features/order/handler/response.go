package handler

type OrderResponse struct {
	ID            uint                `json:"id"`
	UserID        uint                `json:"user_id"`
	TotalPrice    uint                `json:"total_price"`
	Status        string              `json:"status"`
	PaymentMethod string              `json:"payment_method"`
	VANumber      string              `json:"va_number"`
	OrderUniqueID string              `json:"order_unique_id"`
	CreatedAt     string              `json:"created_at"`
	Items         []OrderItemResponse `json:"items"`
}

type OrderItemResponse struct {
	ID           uint   `json:"id"`
	ProductID    uint   `json:"product_id"`
	ProductName  string `json:"product_name"`
	ProductImage string `json:"product_image"`
	ProductPrice uint   `json:"product_price"`
	Quantity     uint   `json:"quantity"`
}