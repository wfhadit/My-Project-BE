package handler

type ProductRequest struct {
	Nama        string `json:"nama" form:"nama" validate:"required"`
	Brand       string `json:"brand" form:"brand" validate:"required"`
	Category    string `json:"category" form:"category" validate:"required"`
	Price       uint   `json:"price" form:"price" validate:"required"`
	Amount      uint   `json:"amount" form:"amount" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	Image       string `json:"image" form:"image" validate:"required"`
}