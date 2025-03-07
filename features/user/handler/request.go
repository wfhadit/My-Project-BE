package handler

type RegisterRequest struct {
	Nama     string `json:"nama" form:"nama" validate:"required,min=3,max=20"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=6,max=20"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=6,max=20"`
}
