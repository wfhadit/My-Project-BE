package handler

type OrderRequest struct {
	TotalPrice    uint               `json:"total_price" form:"total_price" validate:"required"`
	PaymentMethod string             `json:"payment_method" form:"payment_method" validate:"required"`
	Items         []OrderItemRequest `json:"items" form:"items" validate:"required"`
}

type OrderItemRequest struct {
	ProductID    uint   `json:"product_id" form:"product_id" validate:"required"`
	ProductName  string `json:"product_name" form:"product_name" validate:"required"`
	ProductImage string `json:"product_image" form:"product_image" validate:"required"`
	ProductPrice uint   `json:"product_price" form:"product_price" validate:"required"`
	Quantity     uint   `json:"quantity" form:"quantity" validate:"required"`
}