package handler

type LoginResponse struct {
	ID           uint   `json:"id"`
	Nama         string `json:"nama"`
	Email        string `json:"email"`
	TanggalLahir string `json:"tanggal_lahir"`
	JenisKelamin string `json:"jenis_kelamin"`
	NomorHP      string `json:"nomor_hp"`
	Alamat       string `json:"alamat"`
	Foto         string `json:"foto"`
}

type UpdateResponse struct {
	Nama         string `json:"nama"`
	Email        string `json:"email"`
	TanggalLahir string `json:"tanggal_lahir"`
	JenisKelamin string `json:"jenis_kelamin"`
	NomorHP      string `json:"nomor_hp"`
	Alamat       string `json:"alamat"`
	Foto         string `json:"foto"`
}