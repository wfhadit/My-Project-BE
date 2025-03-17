package handler

type CartResponse struct {
	ProductID    uint   `json:"product_id"`
	ProductNama  string `json:"product_nama"`
	ProductImage string `json:"product_image"`
	ProductPrice uint   `json:"product_price"`
	Quantity     uint   `json:"quantity"`
}