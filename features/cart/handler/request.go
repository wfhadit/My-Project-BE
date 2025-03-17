package handler

type CartRequest struct {
	ProductID    uint   `json:"product_id" form:"product_id" `
	ProductNama  string `json:"product_nama" form:"product_nama" `
	ProductImage string `json:"product_image" form:"product_image" `
	ProductPrice uint   `json:"product_price" form:"product_price" `
	Quantity     uint   `json:"quantity" form:"quantity" `
}