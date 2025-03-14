package handler

type ProductResponse struct {
	ID          uint   `json:"id"`
	Nama        string `json:"nama"`
	Brand       string `json:"brand"`
	Category    string `json:"category"`
	Price       uint   `json:"price"`
	Amount      uint   `json:"amount"`
	Description string `json:"description"`
	Image       string `json:"image"`
}