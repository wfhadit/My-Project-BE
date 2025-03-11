package handler

type LoginResponse struct {
	ID    uint   `json:"id"`
	Nama  string `json:"nama"`
	Email string `json:"email"`
	Token string `json:"token"`
}