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

type UpdateRequest struct {
	Nama         string `json:"nama" form:"nama"`
	Email        string `json:"email" form:"email"`
	JenisKelamin string `json:"jenis_kelamin" form:"jenis_kelamin"`
	TanggalLahir string `json:"tanggal_lahir" form:"tanggal_lahir"`
	NomorHP      string `json:"nomor_hp" form:"nomor_hp"`
	Alamat       string `json:"alamat" form:"alamat"`
}
